package main

import "fmt"

func main() {
    var i interface{} = "hello"

    // Get the underlying string value
    s := i.(string)
    fmt.Println(s)

    // A safe way to do it
    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok) // 0 false
}
