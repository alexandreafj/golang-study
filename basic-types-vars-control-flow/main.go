package main

import (
	"fmt"
	"time"
	"unsafe"
)

//Variables and Declarations
var count int
var message string = "Hello"

//Unsigend Intergers
var uin uint = 10
var uin8 uint8 = 255
var uin16 uint16 = 65535
var uin32 uint32 = 4294967295
var uin64 uint64 = 18446744073709551615

// Signed Interegers
var in int = -10
var in8 int8 = -128
var in16 int16 = -32768
var in32 int32 = -2147483648
var in64 int64 = -9223372036854775808

//Floating Point Numbers
var f32 float32 = 3.402823466e+38
var f64 float64 = 1.7976931348623157e+308

//Boolean variables
var isTrue bool = true
var isFalse bool = false

//String Vriable
var str string = "Go Programming"

func main() {
	//Constants Cannot be changed after declaration
	const pi float64 = 3.14159
	const e float64 = 2.71828

	fmt.Println("Pi:", pi)
	fmt.Println("E:", e)
	// ./main.go:44:2: cannot assign to e (neither addressable nor a map index expression)
	//e = 2.718; // This line will cause a compile-time error because constants cannot be reassigned
	fmt.Println("Count:", count)
	fmt.Println("Message:", message)
	fmt.Println("Unsigned Integers:", uin, uin8, uin16, uin32, uin64)
	fmt.Println("Signed Integers:", in, in8, in16, in32, in64)
	fmt.Println("Floating Point Numbers:", f32, f64)
	fmt.Println("Boolean Values:", isTrue, isFalse)
	fmt.Println("String:", str)
	fmt.Println("Size of int:", unsafe.Sizeof(int(0)), "bytes")

	//Control Flow

	// If-Else Statement
	if num := -9; num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Non-negative number")
	}

	//For loops

	//A standart C style for loop
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop Iteration:", i)
	}

	//A while loop equivalent
	i := 0
	for i < 5 {
		fmt.Println("While Loop Iteration:", i)
		i++
	}

	//An inifinite loop
	for {
		fmt.Println("Infinite Loop - Press Ctrl+C to stop")
		break // Break to avoid an actual infinite loop in this example
	}

	//Switch Statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}


	// Compositive Literals
	primes := []int{2,3,5,7,11,13} //Slice literal
	s := make([]int, 5, 10) // make(type, length, capacity)

	fmt.Println("Primes:", primes)
	fmt.Println("Slice s -> len:", len(s), "cap:", cap(s), "data:", s)
}