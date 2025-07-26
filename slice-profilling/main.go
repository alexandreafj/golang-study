package main

// Number of elements to add to our slices.
const numElements = 100000

// createSliceNonPerformant demonstrates the less performant way.
// It starts with a nil slice (len=0, cap=0) and appends to it.
// This will cause the Go runtime to re-allocate the underlying array
// multiple times as the slice grows.
func createSliceNonPerformant() []int {
	var slice []int
	for i := 0; i < numElements; i++ {
		slice = append(slice, i)
	}
	return slice
}

// createSlicePerformant demonstrates the more performant way.
// It uses make() to pre-allocate an underlying array with enough capacity
// for all the elements we plan to add.
// This avoids any re-allocations during the append loop.
func createSlicePerformant() []int {
	slice := make([]int, 0, numElements) // len=0, cap=100000
	for i := 0; i < numElements; i++ {
		slice = append(slice, i)
	}
	return slice
}

// main function to run the code, though we will focus on benchmarks.
func main() {
	_ = createSliceNonPerformant()
	_ = createSlicePerformant()
}