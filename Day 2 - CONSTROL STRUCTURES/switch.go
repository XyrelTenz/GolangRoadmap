package main

import "fmt"

func main() {
    day := "Thursday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week!")
    case "Friday":
        fmt.Println("End of the work week!")
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend! 🎉")
    default:
        fmt.Println("It's a regular day.")
    }
}
