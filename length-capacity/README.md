# Understanding Go Slices: Length vs. Capacity

A quick guide to remember how slice length, capacity, and the slicing syntax work in Go.

---

## The Bus Analogy 🚌

Think of a slice like a bus:

* **Length (`len`)**: The number of people **currently on the bus**.
* **Capacity (`cap`)**: The total number of **seats available on that bus**.

You can add more people (append elements) as long as there are empty seats. When the bus is full (`len == cap`), adding another person requires getting a bigger bus (Go allocates a new, larger underlying array).

---

## Slicing a Slice: The `[start:end]` Rule

The syntax `aSlice[start:end]` creates a new slice from an existing one.

**The rule:** The new slice **includes** the element at `start` but **excludes** the element at `end`.

### Visual Example

Imagine this array and its indices:

```

Index:   0   1   2   3   4
┌───┬───┬───┬───┬───┐───┐───┐
Value: │ 10│ 20│ 30│ 40│ 50│
└───┴───┴───┴───┴───┘───┘───┘

````

When we slice it with `[1:3]`:

1.  **Start at index `1`**: We take the value `20`. ✅
2.  **Continue up to index `3`**: We take the value at index `2`, which is `30`. ✅
3.  **Stop at index `3`**: We have reached the `end` index, so we **stop** and **do not include** the value `40`. 🛑

The result is `[20, 30]`.

### Shorthand

* `aSlice[2:]` means "from index 2 to the very end".
* `aSlice[:3]` means "from the very beginning up to (but not including) index 3".

---

## Full Code Example

This program demonstrates all the concepts in action.

```go
package main

import "fmt"

func main() {
	// 1. Create a slice. The bus has 5 people and 5 seats.
	allSpots := []int{10, 20, 30, 40, 50}
	fmt.Printf("1. allSpots -> len: %d, cap: %d, data: %v\n", len(allSpots), cap(allSpots), allSpots)

	// 2. Create a new slice viewing a part of the original.
	// We are *using* 2 spots (len=2), but our "view" can see 4 total seats (cap=4).
	mySpots := allSpots[1:3] // Elements at index 1 and 2
	fmt.Printf("2. mySpots  -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)

	// 3. Append an element while still within capacity.
	// This uses an empty seat on the *original* bus.
	mySpots = append(mySpots, 99)
	fmt.Printf("3. mySpots (appended) -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)
	fmt.Printf("   allSpots is now changed! -> data: %v\n", allSpots)


	// 4. Append more elements, exceeding capacity.
	// Go gets a bigger bus (a new array), leaving the original behind.
	mySpots = append(mySpots, 88, 77)
	fmt.Printf("4. mySpots (new array) -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)
	fmt.Printf("   allSpots is NOT changed -> data: %v\n", allSpots)
}
````

### Code Output and Explanation

1.  **`allSpots -> len: 5, cap: 5, data: [10 20 30 40 50]`**

      * A slice with 5 elements is created. Its length and capacity are both 5.

2.  **`mySpots -> len: 2, cap: 4, data: [20 30]`**

      * **Length is 2** because we selected 2 elements (`[1:3]` -\> indices 1, 2).
      * **Capacity is 4** because the original array has 4 elements from `mySpots`'s starting point (index 1) to the end: `20, 30, 40, 50`.

3.  **`mySpots (appended) -> len: 3, cap: 4, data: [20 30 99]`**

      * We append `99`. Since `cap` was 4 and `len` was 2, there was room.
      * The `len` becomes 3. The `cap` remains 4.
      * **`allSpots is now changed! -> [10 20 30 99 50]`**: The append operation modified the underlying array that both slices shared.

4.  **`mySpots (new array) -> len: 5, cap: 8, data: [20 30 99 88 77]`**

      * We tried to add 2 elements, but only had 1 spot of capacity left.
      * Go allocated a **brand new, larger array** (with a new capacity, usually double the old one) and copied the elements over.
      * **`allSpots is NOT changed`**: Because `mySpots` now points to a new array, the original `allSpots` is no longer affected.