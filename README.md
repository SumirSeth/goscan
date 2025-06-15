# goscan

`goscan` is a Go module.

## Installation

To use `goscan` in your project, you can use `go get`:

```sh
go get github.com/sumirseth/goscan
```

## Usage

```go
import (
    "fmt"
    "github.com/sumirseth/goscan"
)

func main() {
    // Example usage (replace with actual functionality)
    fmt.Println("Using goscan")
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

---

# TODO:

## Minor Omptimization
- For slightly better memory safety in the future, clone IPs before storing them in a slice:
    ```go
    ips = append(ips, append(net.IP(nil), ip...).String())
    ```