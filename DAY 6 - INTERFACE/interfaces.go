package main

import (
	"fmt"
	"math"
)

// Shape is an interface for things that have an area.
type Shape interface {
    area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

func (c Circle) area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) area() float64 {
    return r.Width * r.Height
}

// measure takes any Shape and prints its area.
func measure(s Shape) {
    fmt.Println("Area:", s.area())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 3}

    measure(c)
    measure(r)
}
