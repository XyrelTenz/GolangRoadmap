package main

import "fmt"

func main() {
    // Declare an array of 3 integers.
    // The length is part of its type.
    var scores [3]int
    scores[0] = 95
    scores[1] = 88
    scores[2] = 92
    fmt.Println("Scores:", scores)
    fmt.Println("Length:", len(scores))
}
