package main

import "fmt"

// Modifying the copied array does not affect the original array because arrays in Go are value types.
// Modifying the copied slice affects the original slice because slices are reference types.
// When creating a slice from an array using `a = b[:]`, modifying the slice impacts the array because slices reference the array.
// Modifying the array after copying the slice does not affect the slice because the slice was copied to the array.

/*
Slice:
a = c creates a reference to the same underlying slice.
a = c[:] also creates a reference to the same underlying slice.
copy(a, c) creates a new slice with copied elements, so changes to a won't affect c.  (shallow copy)

Array:
a = c creates a new array with copied elements, so changes to a won't affect c.  (shallow copy)
a = c[:] creates a new slice with copied elements, so changes to a will affect c.
copy(a, c) When copying a slice, it creates a shallow copy of the elements, meaning it copies the values but doesn't create a new underlying array. For arrays, since their size is fixed, you would generally assign them directly, which creates a copy.
*/

func main() {
	// Array example
	var array1 = [3]int{1, 2, 3} // Declare an array
	var array2 = array1          // Copy the array (shallow copy)
	array2[0] = 10
	fmt.Println("Array1:", array1) // [1 2 3]
	fmt.Println("Array2:", array2) // [10 2 3]

	// Slice example (using a full slice assignment)
	var slice1 = []int{1, 2, 3}
	var slice2 = slice1 // Slice assignment (shallow copy)
	slice2[0] = 10
	fmt.Println("Slice1:", slice1) // [10 2 3]
	fmt.Println("Slice2:", slice2) // [10 2 3]

	// Difference in copying an array vs slice (a = b)
	array3 := [3]int{4, 5, 6}
	slice3 := array3[:] // Create a slice from the array (this is not a copy of the array, but a slice reference)
	slice3[0] = 50
	fmt.Println("Array3:", array3) // [50 5 6]
	fmt.Println("Slice3:", slice3) // [50 5 6]

	// Array and Slice where a is an array and b is a slice
	var slice6 = []int{13, 14, 15}
	var array6 [3]int // Array initialized to zero

	copy(array6[:], slice6)

	array6[0] = 300
	fmt.Println("Array6:", array6) // [300 0 0]
	fmt.Println("Slice6:", slice6) // [13 14 15]

}
