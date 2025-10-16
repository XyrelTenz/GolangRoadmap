package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// area is a method with a receiver of type Rectangle.
func (r Rectangle) area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.area())
}
