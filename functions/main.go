package main

import "fmt"

func combineName(firstName string, lastName string) (string, error) {
	fullname := firstName + " " + lastName
	return fullname, nil
}

func addOne() func() int {
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
	  // this anonymous function
	  // has access to the x variable
	  // defined in the parent function
	  x++
	  return x + 1
	}
}

type Employee struct {
	Name string
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	fullName, err := combineName("Alexandre", "Ferreira")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fullName)

	myFunc := addOne()
	for i := 0; i < 5; i++ {
		fmt.Println(myFunc())
	}

	var employee Employee
    employee.Name = "Alexandre"
    employee.UpdateName("Ferreira")
    employee.PrintName()
}