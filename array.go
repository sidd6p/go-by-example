package main

import "fmt"

func main() {
	// Declaring an integer array with a fixed size of 5 (default values are 0)
	var a [5]int
	fmt.Println("Array a (default values):", a)

	// Assigning a value to the last index of the array
	a[4] = 100
	fmt.Println("Array a (after assignment at index 4):", a)
	fmt.Println("Length of array a:", len(a), cap(a)) // The length of an array is fixed and always known

	// Declaring and initializing an array with specific values
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b (initialized):", b)

	// Declaring and initializing an array with values, using [...] to infer the size
	c := [...]int{1, 2, 3}
	fmt.Println("Array c (size inferred):", c)

	// Declaring an array with explicit indexing during initialization
	// The syntax `3: 6` assigns the value 6 to index 3
	d := [...]int{1, 2, 3: 6, 7}
	fmt.Println("Array d (explicit index assignment):", d)

	// Declaring a 2D array (array of arrays)
	// The outer array has 2 elements, each of which is an array of size 3
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Array e:", e)
}
