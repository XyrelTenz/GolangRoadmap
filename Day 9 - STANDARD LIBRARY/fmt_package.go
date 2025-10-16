package main

import "fmt"

func main() {
    name := "Xyrel"
    age := 20
    fmt.Println("Hello, my name is", name)
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)
    message := fmt.Sprintf("User info: %s, %d", name, age)
    fmt.Println(message)
}
