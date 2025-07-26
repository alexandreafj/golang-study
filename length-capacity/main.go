package main

import "fmt"


func main() {
	// 1. We create an underlying array of 5 parking spots.
	// Go makes a slice pointing to it.
	// The bus has 5 people and 5 seats.
	allSpots := []int{10,20,30,40,50}
	fmt.Printf("allSpots -> len: %d, cap: %d, data: %v\n", len(allSpots), cap(allSpots), allSpots)
	fmt.Println("--------------------------------------------------")

	// 2. We create a new slice `mySpots` that only looks at spots 20 and 30.
	// We're only *using* 2 spots, so length is 2.
	// But our "view" can still see the original spots 40 and 50.
	// So, we have 2 people, but 4 total seats available from our starting point.
	mySpots := allSpots[1:3] // Elements at index 1 up to (but not including) 3
	fmt.Printf("mySpots  -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)
	fmt.Println("--------------------------------------------------")


	// 3. Let's add a car to spot 99. We have empty seats (capacity), so we can.
	// We add one person, so length becomes 3. Capacity is still 4.
	// Notice this changes the *original* `allSpots` array because we are using its capacity.
	mySpots = append(mySpots, 99)
	fmt.Printf("mySpots  -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)
	fmt.Printf("allSpots is now changed! -> data: %v\n", allSpots)
	fmt.Printf("allSpots -> len: %d, cap: %d\n", len(allSpots), cap(allSpots))
	fmt.Println("--------------------------------------------------")

	// 4. Now, let's add two more cars. Our bus only has 1 empty seat left! (len 3, cap 4)
	// Go must get a bigger bus! It creates a brand-new, larger array.
	// The capacity will now be much larger, and it's no longer connected to `allSpots`.
	mySpots = append(mySpots, 88, 77)
	fmt.Printf("mySpots  -> len: %d, cap: %d, data: %v\n", len(mySpots), cap(mySpots), mySpots)
	fmt.Printf("allSpots is NOT changed now -> data: %v\n", allSpots)
	fmt.Printf("allSpots -> len: %d, cap: %d", len(allSpots), cap(allSpots))
}