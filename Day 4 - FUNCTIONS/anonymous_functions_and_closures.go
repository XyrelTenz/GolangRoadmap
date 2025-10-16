package main

import "fmt"

// intSeq returns a function (a closure).
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    // Anonymous function
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println("Anonymous add:", add(2, 2))

    // Closure example
    nextInt := intSeq()
    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
}
