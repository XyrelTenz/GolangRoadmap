package main

import "fmt"

// describe takes an empty interface and can handle any type.
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
    describe(42)
    describe("hello")
    describe(true)
}
