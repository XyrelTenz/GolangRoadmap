package Playground

func Closures() func() int {
	// Local Variable
	count := 0

	// Closure Function
	return func() int {
		//Captures the local variable count
		count++
		return count
	}
}
