package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func ExpandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}
	// Remove network and broadcast addresses if needed
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

func ExpandPorts(ports string) ([]string, error) {
	var result []string
	items := strings.Split(ports, ",")

	for _, item := range items {
		if strings.Contains(item, "-") {
			parts := strings.Split(item, "-")
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil || start > end {
				return nil, fmt.Errorf("invalid port range: %s", item)
			}
			for p := start; p <= end; p++ {
				result = append(result, strconv.Itoa(p))
			}
		} else {
			result = append(result, item)
		}
	}
	return result, nil
}


// Helper function to increment an IP address
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func scanPorts(ips []string, ports string) {
	portList, err := ExpandPorts(ports)
	if err != nil {
		fmt.Println("Port parse error:", err)
		return
	}

	var wg sync.WaitGroup
	for _, ip := range ips {
		for _, port := range portList {
			wg.Add(1)
			go func(ip, port string) {
				defer wg.Done()
				fmt.Printf("Scanning %s:%s\n", ip, port)
			}(ip, port)
		}
	}
	wg.Wait()
}


func main() {
	// go run main.go -ip 192.168.1.0/30 -ports 22,80,443,8000-8002
	// reference above
	// -ip and -ports are required
	// -ports can be a single port or a range of ports
	var rootCmd = &cobra.Command{
		Use:   "goscan",
		Short: "A fast, concurrent TCP port scanner and HTTP service mapper.",
		Long:  "goscan: A fast, concurrent TCP port scanner for IPs/CIDRs. Identifies open ports and analyzes HTTP/S services (status, headers, title), with JSON/CSV output.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ip, err := cmd.Flags().GetString("ip")
			if err != nil {
				return err
			}
			ports, err := cmd.Flags().GetString("ports")
			if err != nil {
				return err
			}

			ips, err := ExpandCIDR(ip)
			if err != nil {
				return err
			}

			scanPorts(ips, ports)

			return nil
		},
	}

	var ip string

	rootCmd.Flags().StringVar(&ip, "ip", "", "The IP address or CIDR range to scan.")
	rootCmd.MarkFlagRequired("ip")

	var ports string

	rootCmd.Flags().StringVar(&ports, "ports", "", "The ports to scan. Can be a single port or a range of ports.")
	rootCmd.MarkFlagRequired("ports")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}