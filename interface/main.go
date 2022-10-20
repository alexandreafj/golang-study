package main

import "fmt"

type Employee interface{
	GetName() string
}

type Engineer struct {
	name string
}

func(e *Engineer) GetName() string {
	return e.name;
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	engineer := Engineer{name: "Alexandre"}
	PrintDetails(&engineer);
}