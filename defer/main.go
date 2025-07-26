package main

import "fmt"

func main() {
    fmt.Println("start")

		// LIFO (Last In, First Out) order
		defer fmt.Println("four")  // This is deferred first
    defer fmt.Println("one")   // This is deferred last
    defer fmt.Println("two")   // This is deferred second
    defer fmt.Println("three") // This is deferred third

    fmt.Println("finish")
}