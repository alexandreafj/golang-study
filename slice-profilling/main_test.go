package main

import "testing"

// Benchmark for the non-performant version.
// Go will run this function multiple times to get a reliable average.
func BenchmarkCreateSliceNonPerformant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createSliceNonPerformant()
	}
}

// Benchmark for the performant version.
func BenchmarkCreateSlicePerformant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createSlicePerformant()
	}
}
