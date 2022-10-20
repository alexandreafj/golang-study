package main

import (
	"testing"
)

// How to test
// go test
// go test -v verbose
// go test -cover coverage of the code
// Shows where in the code it's not cover by
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// Benchmark a test go test -bench=.
// you need to type "Benchmark" first as the name of a function

func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        if output := Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}

func benchmarkCalculate(input int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        Calculate(input)
    }
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }

