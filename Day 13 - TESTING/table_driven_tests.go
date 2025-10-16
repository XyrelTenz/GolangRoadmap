package main

import "testing"

func TestAddTableDriven(t *testing.T) {
    var tests = []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
    }

    for _, test := range tests {
        if output := Add(test.a, test.b); output != test.expected {
            t.Errorf("Test Failed: %d + %d expected %d, got %d", test.a, test.b, test.expected, output)
        }
    }
}
