package main

import "fmt"

func main() {
    // Maps are Go's hash table / dictionary type.
    // map[key-type]value-type
    studentAges := make(map[string]int)

    studentAges["John"] = 22
    studentAges["Jane"] = 21

    fmt.Println("Student Ages:", studentAges)
    fmt.Println("Jane's age:", studentAges["Jane"])

    // Delete a key
    delete(studentAges, "John")
    fmt.Println("Map after deletion:", studentAges)
}
