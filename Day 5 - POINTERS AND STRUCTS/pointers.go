package main

import "fmt"

// zeroval doesn't change the original variable.
func zeroval(ival int) {
    ival = 0
}

// zeroptr changes the original variable via a pointer.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i) // Pass the memory address of i
    fmt.Println("zeroptr:", i)
}
