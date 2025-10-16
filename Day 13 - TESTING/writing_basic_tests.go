package main

import "testing"

func Add(x, y int) int {
	return x + y
}
func TestAdd(t *testing.T) {
    sum := Add(2, 2)
    expected := 4
    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}
