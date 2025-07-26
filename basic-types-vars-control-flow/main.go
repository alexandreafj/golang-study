package main

import (
	"fmt"
	"time"
	"unsafe"

	rectangle "github.com/alexandreafj/golang-study/basic-types/Rectangle"
	// Importing the rectangle package
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

//Structs
type Person struct {
	Name string
	Age  int
}

type Point struct {
	X, Y int
}

func ChangeSlice(s *[]int) {
	// This function modifies the slice passed to it
	(*s)[0] = 100
	(*s)[1] = 200
	(*s)[2] = 300
	(*s)[3] = 400
	(*s)[4] = 500
	*s = append(*s, 600)
	fmt.Println("Inside ChangeSlice -> len:", len(*s), "cap:", cap(*s), "data:", *s)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

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

	// Change the slice appending new values without using pointers
	// ChangeSlice(&s) // Passing the address of the slice to modify it
	// If you don't want to use pointers, when it finishes the function the changes will not be reflected in the original slice
	ChangeSlice(&s)
	fmt.Println("Slice s -> len:", len(s), "cap:", cap(s), "data:", s)

	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println("Ages Map:", ages)

	age,ok := ages["charlie"]
	if!ok {
		fmt.Println("Charlie not found in ages map")
	} else {
		fmt.Println("Charlie's age:", age)
	}

	p := Person{Name: "John", Age: 30}
	fmt.Println("Person:", p)

	reult, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", reult)
	}

	rect := rectangle.Rectangle{Width: 5, Height: 10}
	fmt.Println("Rectangle Area:", rect.Area())


	//Pointers
	var ptr *int
	var numPtr *int = &count // Pointer to count variable
	fmt.Println("Pointer to count:", numPtr, "Value:", *numPtr)
	fmt.Println("Size of pointer:", unsafe.Sizeof(ptr), "bytes")

	var point *Point = &Point{X: 10, Y: 20}
	fmt.Println("Point:", point, "X:", point.X, "Y:", point.Y)
	fmt.Println("Size of pointer:", unsafe.Sizeof(point), "bytes")
}