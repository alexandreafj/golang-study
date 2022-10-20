package main

import "fmt"

// all numeric types default to 0

// unsigned int with 8 bits
// Can store: 0 to 255
var myintu8 uint8
// signed int with 8 bits
// Can store: -127 to 127
var myint8 int8

// unsigned int with 16 bits
var myintu16 uint16
// signed int with 16 bits
var myint16 int16

// unsigned int with 32 bits
var myintu23 uint32
// signed int with 32 bits
var myint32 int32

// unsigned int with 64 bits
var myintu64 uint64
// signed int with 64 bits
var myint64 int64

type Person struct {
	name string
	age int
}

type Team struct {
	name string
	players [2]Person
}

func main() {
	fmt.Println("Hello World")
	// cannot use 2500 (untyped int constant) as int8 value in variable declaration (overflows)compilerNumericOverflow
	//var explodeInt int8 = 2500
    var myint int8
    for i := 0; i < 2500; i++ {
        myint += 1
    }
    fmt.Println(myint) // prints out -127


	var men uint8 = 5
	var women int16 = 6
	// invalid operation: men + women (mismatched types uint8 and int16)compilerMismatchedTypes
	//var people int = men + women
	var people int = int(men) + int(women)
	fmt.Println(people);

	// Floating point numbers
	// var f1 float32
	// var f2 float64
	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you think this will return false but it returns true
	fmt.Println(maxFloat32+10) // 16777216
	fmt.Println(maxFloat32+2000000) // 16777216


	// converting from int to float
	// var myintconvert int
	// myfloat := float64(myintconvert)

	// converting from float to int
	// var myfloat2convert float64
	// myint2 := int(myfloat2convert)

	// complex64 = float32 * 2
	// complex128 = float64 * 2
	var a complex64 = complex(1, 2)
	var b complex64 = 2 + 4i
	var c complex64 = b - a
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var isTrue bool = true
	var isFalse bool = false
	// AND
	if isTrue && isFalse {
	  fmt.Println("Both Conditions need to be True")
	}
	// OR - not exclusive
	if isTrue || isFalse {
	  fmt.Println("Only one condition needs to be True")
	}

	// const is a immutable variable
	const meaningOfLife = "I dont know";
	//cannot assign to meaningOfLife (untyped string constant "I dont know")compilerUnassignableOperand
	// meaningOfLife = 1;

	// var days []string
	days := [...]string{"monday", "tuesday", "wednesday", "thurday", "friday", "staurday", "sunday"}
	//slice
	fmt.Println(days[0]);
	fmt.Println(days[5]);

	weekdays := days[0:5]
	fmt.Println(weekdays)

	youtubeSubs := map[string]int {
		"urn:teste:teste": 1000,
		"urn:abc:abc": 100,
	}

	fmt.Println(youtubeSubs["urn:teste:teste"])

	elliot := Person{name: "Elliot", age: 24}

	elliot.age = 18

	fmt.Println(elliot)

	players := [...]Person{{name: "Forrest", age: 18}, {name: "Bruce", age: 20}}
	celtic := Team{name: "Celtic FC", players: players}

	fmt.Printf("%+v\n", celtic)
}

