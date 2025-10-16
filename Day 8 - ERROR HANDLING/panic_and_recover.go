package main

import "fmt"

func mayPanic() {
    // A function that might panic
    panic("a problem!")
}

func main() {
    // defer runs after the surrounding function completes.
    // recover can stop a panic and resume execution.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    mayPanic()

    // This line will not be reached because of the panic.
    fmt.Println("After mayPanic")
}
