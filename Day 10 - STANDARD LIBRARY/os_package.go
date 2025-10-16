// Run this from your terminal: go run 01_os_package.go arg1 arg2
package main

import (
	"fmt"
	"os"
)

func main() {
    // os.Args provides access to command-line arguments.
    // The first argument is the path of the program itself.
    args := os.Args
    fmt.Println("All arguments:", args)
    fmt.Println("Program name:", args[0])
    if len(args) > 1 {
        fmt.Println("First argument:", args[1])
    }
}
