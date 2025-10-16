package main

import (
	"fmt"

	"roadmapgolang/day1"
)

func main() {
	calc := day1.Calculators{A: 5, B: 3}
	fmt.Println(calc.Add())
}
