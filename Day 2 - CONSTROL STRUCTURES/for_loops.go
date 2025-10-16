package main

import "fmt"

func main() {
    // Classic for loop
    for i := 0; i < 3; i++ {
        fmt.Println("Classic loop:", i)
    }

    // For loop as a "while" loop
    sum := 1
    for sum < 5 {
        sum += sum
    }
    fmt.Println("Sum (while-style):", sum)

    // For loop over a slice (range)
    nums := []int{10, 20, 30}
    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
