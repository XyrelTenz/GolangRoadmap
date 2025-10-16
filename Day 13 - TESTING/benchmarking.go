package main

import "testing"

// Add function needs to be here if not in a separate file
func Add(x, y int) int {
    return x + y
}

func BenchmarkAdd(b *testing.B) {
    // b.N is a number of iterations chosen by the testing framework
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
