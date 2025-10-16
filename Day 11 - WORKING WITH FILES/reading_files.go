package main

import (
	"fmt"
	"os"
)

func main() {
    // Create a dummy file first
    _ = os.WriteFile("test.txt", []byte("Hello, Go!"), 0666)

    // Read the entire file content
    content, err := os.ReadFile("test.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println("File content:", string(content))
}
