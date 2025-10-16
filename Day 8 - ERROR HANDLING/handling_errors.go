package main

import (
	"fmt"
	"strconv"
)

func main() {
    // The common 'if err != nil' pattern
    numStr := "123a"
    num, err := strconv.Atoi(numStr)
    if err != nil {
        fmt.Println("Could not convert string to int.")
        fmt.Println("Error message:", err)
        return // Stop execution
    }
    fmt.Println("Converted number:", num)
}
