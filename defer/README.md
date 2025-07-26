
# The `defer` Keyword in Go

The `defer` keyword in Go schedules a function call to be executed immediately before the function it's in returns. It's a powerful tool for simplifying code, especially for cleanup actions like closing files or unlocking mutexes.

-----

## How It Works

`defer` operates on two main principles:

1.  **Delayed Execution**: The function call isn't executed right away. Instead, it's placed onto a special stack associated with the surrounding function.
2.  **LIFO Order (Last-In, First-Out)**: If multiple `defer` statements are used, they are executed in the reverse order of their declaration. The last deferred call is the first one to be executed.

### An Important Detail: Argument Evaluation

When a function call is deferred, its **arguments are evaluated immediately**, but the **function call itself is delayed**. The evaluated arguments are stored and used later when the deferred call is executed.

Consider this example:

```go
package main

import "fmt"

func main() {
    value := 1

    // The argument 'value' is evaluated here. Its value is 1.
    defer fmt.Println("Deferred print:", value)

    value = 2
    fmt.Println("Current value:", value)
}
```

**Output:**

```
Current value: 2
Deferred print: 1
```

Even though `value` was `2` at the end of the function, the deferred call printed `1` because that was its value when the `defer` statement was executed.

-----

## Example: Guaranteed File Closing

The most common use case for `defer` is to guarantee that resources are released. By deferring the cleanup action right after the resource is acquired, you make it impossible to forget.

```go
package main

import (
	"fmt"
	"os"
)

func writeFile(filename, text string) error {
	// 1. A file is opened, acquiring a resource from the OS.
	file, err := os.Create(filename)
	if err != nil {
		// If we fail here, the function returns, and nothing is deferred.
		return err
	}
	// 2. We immediately defer the Close() call. This guarantees that
	//    the file will be closed when this function (writeFile) returns,
	//    no matter what happens next.
	defer file.Close()

	// 3. We perform our work with the file.
	_, err = file.WriteString(text)
	if err != nil {
		// Even if we return early due to an error here, the deferred
		// file.Close() will still run before the function exits.
		return err
	}

	fmt.Printf("Successfully wrote to %s\n", filename)
	return nil // The deferred file.Close() runs right before this return.
}

func main() {
	err := writeFile("hello.txt", "Hello, defer!")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
```

This pattern makes the code robust and easy to read because the cleanup logic is placed right next to the allocation logic.