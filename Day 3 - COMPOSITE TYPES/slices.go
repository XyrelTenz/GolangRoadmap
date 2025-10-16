package main

import "fmt"

func main() {
    // Slices are more common and flexible than arrays.
    grades := []int{95, 88, 92}
    fmt.Println("Initial grades:", grades)

    // Add elements using append
    grades = append(grades, 78)
    fmt.Println("Appended grades:", grades)

    // Slicing a slice
    topTwo := grades[0:2]
    fmt.Println("Top two grades:", topTwo)
}
