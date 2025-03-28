# go-time

A Go library for parsing time duration strings into milliseconds. It provides a simple and intuitive way to convert human-readable time strings into their millisecond equivalents.

## Features

- Convert time strings to milliseconds
- Support for multiple time units
- Simple and easy-to-use API
- Zero external dependencies

## Installation

```bash
go get github.com/iiivanpopov/go-time
```

## Supported Time Units

- `w` - weeks
- `d` - days
- `h` - hours
- `m` - minutes
- `s` - seconds
- `S` - milliseconds

## Usage

```go
package main

import (
    "fmt"
    "github.com/iiivanpopov/go-time/gotime"
)

func main() {
    // Convert hours to milliseconds
    millis, err := gotime.ToMillis("2h")
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("2 hours in milliseconds:", millis) // Output: 7200000

    // Convert days to milliseconds
    millis, err = gotime.ToMillis("1d")
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("1 day in milliseconds:", millis) // Output: 86400000

    // Convert weeks to milliseconds
    millis, err = gotime.ToMillis("1w")
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("1 week in milliseconds:", millis) // Output: 604800000
}
```

## Error Handling

The library provides several error types for proper error handling:

- `ErrInvalidLength`: Time string must have at least 2 characters
- `ErrNegativeValue`: Negative time values are not allowed
- `ErrInvalidUnit`: Invalid time unit was provided
- `ErrInvalidFormat`: Invalid number format in the time string

## License

This project is available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
