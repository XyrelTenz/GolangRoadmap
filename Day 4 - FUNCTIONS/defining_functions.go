package main

import "fmt"

// add takes two integers and returns their sum.
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("5 + 3 =", result)
}
