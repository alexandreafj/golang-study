package main

import "fmt"

type Service interface{
	SayHi()
}

type MyService struct{}
func (s MyService) SayHi() {
	fmt.Println("Hi")
}

type SecondService struct{}
func (s SecondService) SayHi() {
	fmt.Println("Hello From the 2nd Service")
}

func main() {
	mymap := make(map[string]int)

	mymap["teste"] = 10
	mymap["teste"] = 11
	mymap["teste"] = 12

	for key, value := range mymap {
		fmt.Println(key)
		fmt.Println(value)
	}

	interfaceMap := make(map[string]Service)
	
	interfaceMap["SERVICE-ID-1"] = MyService{}
	interfaceMap["SERVICE-ID-2"] = SecondService{}

	interfaceMap["SERVICE-ID-2"].SayHi()
}