package main

import (
  	"fmt"
  	"os"

  	"github.com/spf13/cobra"
)

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
			var ip = cmd.Flags().Lookup("ip").Value.String()
			if ip == "" {
				return cmd.Help()
			}
			var ports = cmd.Flags().Lookup("ports").Value.String()
			

			fmt.Println("IP:", ip)
			fmt.Println("Ports:", ports)

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