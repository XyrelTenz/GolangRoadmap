package day1

type Calculators struct {
	A int
	B int
}

func (c Calculators) Add() int {
	return c.A + c.B
}
