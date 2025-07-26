
The saying **"Don't communicate by sharing memory; share memory by communicating"** is a guideline for handling data in concurrent programs.

  * **"Don't communicate by sharing memory"** refers to the traditional method where multiple threads (or goroutines) access the same piece of memory. To prevent data corruption (a "race condition"), you must use locks (like a `sync.Mutex`) to ensure only one goroutine can access that memory at a time. This method can be complex and error-prone.

  * **"Share memory by communicating"** is the idiomatic Go approach. Instead of goroutines fighting over access to a shared variable, you pass the data from one goroutine to another through a **channel**. This way, only one goroutine has ownership of the data at any given moment, which naturally prevents race conditions. The communication itself is the mechanism that transfers ownership.

In short: **Don't use locks to protect shared data; use channels to pass copies of data between goroutines.**

-----

### The "Old Way": Communicating by Sharing Memory (Using Mutex)

This is the pattern the Go proverb advises against. Here, multiple goroutines try to increment the same `counter` variable. To do this safely, we must wrap every access to `counter` with a `mutex.Lock()` and `mutex.Unlock()`.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex // A mutex (lock) to protect the counter

	// We want to run 1000 operations concurrently
	numOperations := 1000
	wg.Add(numOperations)

	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			
			// Lock the memory so no other goroutine can access it
			mu.Lock()
			// Modify the shared memory
			counter++
			// Unlock the memory so others can use it
			mu.Unlock()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter (with mutex):", counter) // Correctly prints 1000
}
```

This works, but it has downsides:

  * You have to remember to lock and unlock everywhere you use the variable.
  * Forgetting to unlock can freeze your program (deadlock).
  * It's harder to reason about the flow of data.

-----

### The "Go Way": Share Memory by Communicating (Using Channels)

Here's the same problem solved using a channel. Instead of all goroutines modifying a shared `counter`, each goroutine does its small piece of work and sends the result over a channel. The main goroutine is the sole owner of the `counter` and is the only one that modifies it.

A **channel** is a typed conduit that allows you to send and receive values between goroutines. Think of it as a pipe. ⚙️

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	
	// Create a channel to receive results from our goroutines.
	// This channel can only transport integers.
	results := make(chan int)

	numOperations := 1000
	wg.Add(numOperations)

	// Launch the worker goroutines
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			// Instead of modifying a shared variable,
			// send the result of your work (the number 1) over the channel.
			results <- 1
		}()
	}

	// Start a separate goroutine to close the channel once all workers are done.
	// This is crucial to prevent the main goroutine from waiting forever.
	go func() {
		wg.Wait()
		close(results)
	}()

	// The main goroutine now safely receives all the results from the channel
	// and aggregates them. No mutex needed!
	for val := range results {
		counter += val
	}

	fmt.Println("Final counter (with channel):", counter) // Correctly prints 1000
}
```

In this version, no two goroutines are accessing the same data. The data (the number `1`) is copied and passed through the channel. This makes the logic much cleaner and safer.

-----

## How Channels Work

Channels are the heart of Go concurrency. Here are the key operations:

### Creation

You create a channel using the `make()` function.

  * **Unbuffered Channel:** `ch := make(chan int)`

      * A send operation (`ch <- val`) on an unbuffered channel will **block** (pause the goroutine) until another goroutine is ready to receive the value (`<-ch`).
      * This forces synchronization between the sender and receiver.

  * **Buffered Channel:** `ch := make(chan string, 5)`

      * This channel has a buffer and can hold 5 values.
      * A send operation will only block if the buffer is full. A receive will only block if the buffer is empty.
      * This decouples the sender and receiver, as they don't have to be ready at the exact same time.

### Sending and Receiving

The `<-` operator is used for both sending and receiving.

```go
ch := make(chan string)

// Send a value into the channel (in a separate goroutine)
go func() {
	ch <- "Hello, channel!"
}()

// Receive a value from the channel
message := <-ch 
fmt.Println(message) // Prints "Hello, channel!"
```

### Closing and Ranging

When you're done sending values, you can `close()` a channel. This signals to receivers that no more data is coming.

You can easily read all values from a channel until it's closed using a `for range` loop.

```go
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	
	// It's important to close a channel if you plan to range over it.
	// Otherwise, the loop will wait forever for more values (deadlock).
	close(queue)

	// This loop receives values from `queue` until it's empty and closed.
	for elem := range queue {
		fmt.Println(elem)
	}
}
// Output:
// one
// two
```