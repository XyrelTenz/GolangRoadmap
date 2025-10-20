package main

import (
	"fmt"
)

func main() {

	// miniProjects.BankSystem(&miniProjects.BankStruct{})
	// DSA.Fibonnaci()
	// Playground.Validation()

	defer func() {
		fmt.Println("Second")
	}()
	fmt.Println("First")

}
