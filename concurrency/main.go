package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayWihtoutWG(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go say("world", &wg)
	sayWihtoutWG("hello")

	// Wait for all goroutines that we Added to the group to call Done()
  fmt.Println("Main goroutine is waiting for other goroutines to finish...")
  wg.Wait()

  fmt.Println("All goroutines finished. Main goroutine is exiting.")
}