package main

import "fmt"

/*
When you create a slice from another slice, the new slice references the same underlying slice.
Modifying the elements of the new slice will affect the original slice and the underlying slice.
*/

/*
Arrays in Go are value types, meaning when you assign an array to another array or pass an array to a function,
a copy of the array is created, not a reference to the original array.
As a result, modifying the new array will not affect the original array.
In contrast, slices are reference types, so they behave differently in terms of modifying the underlying data.
*/

func main() {
	// Declaring an empty slice of integers (nil slice)
	var a []int
	fmt.Println("Slice a (nil slice):", a, "| Is nil:", a == nil, "| Length:", len(a), "| Capacity:", cap(a))

	// Declaring and initializing a slice with values
	var b = []int{1, 2, 3, 4}
	fmt.Println("Slice b (initialized):", b, "| Length:", len(b), "| Capacity:", cap(b))

	// Creating a slice of strings with a predefined length of 3 (elements are zero-valued)
	// Capacity: Maximum number of elements the slice can hold before it needs to grow is set to 10 here
	s := make([]string, 3, 10)
	fmt.Println("Slice s (created with make):", s, "| Is nil:", s == nil, "| Length:", len(s), "| Capacity:", cap(s))

	// Assigning values to the slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Slice s (after assigning values):", s)

	// Issue with assigning an element out of bounds (index 3 is invalid for a slice of length 3):
	// s[3] = "d" // This would panic: index out of range

	// Appending values to the slice (dynamic resizing)
	// append cannot be used with arrays as they are fixed in size
	s = append(s, "d")      // Adding one element
	s = append(s, "e", "f") // Adding multiple elements
	fmt.Println("Slice s (after appending values):", s)

	// Creating a new slice c with the same length as s and copying values
	c := make([]string, len(s))
	fmt.Println("Slice c (created for copying):", c, "| Length:", len(c), "| Capacity:", cap(c))
	copy(c, s) // Copying values from s to c
	fmt.Println("Slice c (after copying values):", c)
}
