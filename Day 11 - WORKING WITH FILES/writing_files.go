package main

import (
	"fmt"
	"os"
)

func main() {
    data := []byte("This is some new data.")
    // WriteFile creates the file if it doesn't exist,
    // or truncates it before writing if it does.
    err := os.WriteFile("output.txt", data, 0666)
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully wrote to output.txt")
}
