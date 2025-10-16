package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
    // Create an io.Reader from a string.
    r := strings.NewReader("Hello, Reader!")

    // Read from it in chunks.
    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
}
