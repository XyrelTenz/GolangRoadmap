package main

import "fmt"

func main() {
    // Using 'var' to declare a variable
    var name string = "Go Developer"
    fmt.Println(name)

    // Using short declaration operator ':='
    // Go infers the type
    level := 1
    fmt.Println("Current Level:", level)

    // Declaring a constant
    const pi = 3.14159
    fmt.Println("Pi value:", pi)
}
