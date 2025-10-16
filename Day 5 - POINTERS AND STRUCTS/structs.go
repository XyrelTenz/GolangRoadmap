package main

import "fmt"

// A Person struct has name and age fields.
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a new struct instance
    p1 := Person{Name: "Alice", Age: 30}
    fmt.Println(p1)
    fmt.Println("Name:", p1.Name)
}
