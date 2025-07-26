package main

import (
	"fmt"
	"time"
)

// func main() {
// 	var counter int
// 	var wg sync.WaitGroup

// 	// Create a channel to receive results from our goroutines.
// 	// This channel can only transport integers.
// 	results := make(chan int)

// 	numOperations := 1000
// 	wg.Add(numOperations)

// 	// Launch the worker goroutines
// 	for i := 0; i < numOperations; i++ {
// 		go func() {
// 			defer wg.Done()
// 			// Instead of modifying a shared variable,
// 			// send the result of your work (the number 1) over the channel.
// 			results <- 1
// 		}()
// 	}

// 	// Start a separate goroutine to close the channel once all workers are done.
// 	// This is crucial to prevent the main goroutine from waiting forever.
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// The main goroutine now safely receives all the results from the channel
// 	// and aggregates them. No mutex needed!
// 	for val := range results {
// 		counter += val
// 	}

// 	fmt.Println("Final counter (with channel):", counter) // Correctly prints 1000
// }

func main() {
	ch := make(chan string)

	go func() {
		// This operation takes 2 seconds to complete
		time.Sleep(1 * time.Microsecond)
		ch <- "operation complete"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // We only wait for 1 second
		fmt.Println("Error: operation timed out ⏳")
	}
}