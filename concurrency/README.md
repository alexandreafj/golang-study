-----

# Concurrency in Go: Goroutines and WaitGroup

Go offers powerful built-in features for concurrency, enabling you to execute tasks simultaneously. The fundamental unit of concurrency in Go is the **goroutine**.

-----

## What is a Goroutine?

A goroutine is a lightweight thread managed by the Go runtime. They are significantly more efficient than traditional operating system threads, allowing you to run thousands or even millions of them concurrently.

To initiate a new goroutine, simply use the `go` keyword before a function call. The Go runtime will then execute this function concurrently with the rest of your program.

```go
// This function will run in its own goroutine
go myFunction()

// The main program continues immediately without waiting for myFunction() to finish
...
```

-----

## The Main Goroutine

Every Go program begins with a single goroutine, known as the **main goroutine**, which executes the `main()` function.

A crucial rule in Go is that **the program exits as soon as the main goroutine finishes**. It does not wait for other goroutines to complete their work. This can lead to a common issue where your program terminates before other goroutines have a chance to run.

-----

## Synchronizing Goroutines with `sync.WaitGroup`

To address the problem of the main goroutine exiting prematurely, we need a mechanism to make it wait for other goroutines to complete. The standard approach for this in Go is using `sync.WaitGroup`.

A `WaitGroup` acts as a counter that blocks the execution of a goroutine until its internal counter reaches zero. It provides three primary methods:

  * `wg.Add(n)`: Increments the counter by `n`. You call this before starting the goroutine(s) to signal how many tasks you are about to launch.
  * `wg.Done()`: Decrements the counter by one. Each goroutine *must* call this when it finishes its task. It's common practice to use `defer wg.Done()` to ensure it's always called.
  * `wg.Wait()`: Blocks the goroutine where it's called (typically the main goroutine) until the counter becomes zero.

-----

## Example: Waiting for Goroutines to Finish

Here's a complete example demonstrating how to use a `WaitGroup` to manage multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// This function will be run by each of our goroutines.
// It takes a WaitGroup pointer so it can signal when it's done.
func worker(id int, wg *sync.WaitGroup) {
	// Defer wg.Done() to ensure the counter is decremented when the function exits.
	// This is a crucial step!
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work
	time.Sleep(time.Second)

	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	// 1. Create a WaitGroup.
	var wg sync.WaitGroup

	// 2. Launch several goroutines.
	for i := 1; i <= 3; i++ {
		// Increment the WaitGroup counter for each goroutine we are about to start.
		wg.Add(1)

		// Launch a new goroutine, passing the worker ID and a pointer to the WaitGroup.
		go worker(i, &wg)
	}

	fmt.Println("Main: Waiting for workers to finish...")

	// 3. Block the main goroutine until the WaitGroup counter is zero.
	//    The program will pause here until all workers have called wg.Done().
	wg.Wait()

	fmt.Println("Main: All workers have finished. Exiting.")
}
```

-----

### Execution Flow:

1.  The `main` function creates a `WaitGroup`.
2.  It then loops three times. In each iteration, it calls `wg.Add(1)` (incrementing the counter) and subsequently launches a `worker` goroutine.
3.  The `main` goroutine continues and reaches `wg.Wait()`. At this point, the counter is 3, so `main` blocks and waits.
4.  Meanwhile, the three `worker` goroutines are running concurrently. Each one prints its start message, simulates work by sleeping for a second, prints its finish message, and then calls `wg.Done()`, decrementing the counter.
5.  Once the third and final worker calls `wg.Done()`, the counter becomes zero.
6.  The `wg.Wait()` in `main` unblocks, and the program prints the final message before exiting cleanly.