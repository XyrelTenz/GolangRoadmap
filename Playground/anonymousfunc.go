package Playground

import "fmt"

func anonymousFunc() {
	anonymous := func() int {
		fmt.Println("Anonymous function")

		return 0

	}

	fmt.Print(anonymous())
}
