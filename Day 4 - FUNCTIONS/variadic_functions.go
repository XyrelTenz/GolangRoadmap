package main

import "fmt"

// sum accepts a variable number of integers.
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum(1, 2):", sum(1, 2))
    fmt.Println("Sum(1, 2, 3, 4):", sum(1, 2, 3, 4))
}
