package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(" = ", total)
}

func Calculator(num1 int, num2 int) (calculate int) {
	fmt.Print("First Number: ")
	fmt.Scanln(&num1)

	fmt.Print("Second Number: ")
	fmt.Scanln(&num2)

	calculate = num1 + num2
	return calculate
}

func main() {

	// miniProjects.BankSystem(&miniProjects.BankStruct{})
	// DSA.Fibonnaci()
	// Playground.Validation()

	sum(1, 2, 3, 4, 5)
	Calculator(2, 3)
}
