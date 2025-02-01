package main

import "fmt"

/*
Method: Tied to a specific instance of rect, called with . (e.g., r.area()).
Function: Takes rect as an argument, called with the function name (e.g., area(r)
*/

// Define a struct `rect` representing a rectangle with `width` and `length` fields.
type rect struct {
	width, length int
}

// Define a method `area` to calculate the area of the rectangle.
// This method has a receiver of type `*rect` (pointer to `rect`).
// Using a pointer receiver allows the method to modify the original struct if needed
// and avoids copying the struct on method calls.
func (r *rect) area() int {
	return r.length * r.width
}

// Define a method `preim` to calculate the perimeter of the rectangle.
// This method has a receiver of type `rect` (value type).
// Using a value receiver means the method operates on a copy of the struct,
// so it cannot modify the original struct.
func (r rect) preim() int {
	return 2*r.width + 2*r.length
}

func main() {
	// Create an instance of `rect` with width 10 and length 5.
	r := rect{10, 5}

	// Call methods on the `rect` instance.
	// `area` (pointer receiver) is called on a value (`r`), but Go automatically handles
	// the conversion to a pointer.
	// `preim` (value receiver) is called directly on the value (`r`).
	fmt.Println(r.area(), r.preim()) // Output: 50 30

	// Create a pointer to the `rect` instance.
	rp := &r

	// Call methods on the pointer to the `rect` instance.
	// `area` (pointer receiver) works directly with the pointer.
	// `preim` (value receiver) can also be called on a pointer because Go automatically
	// dereferences the pointer to call the method on the value.
	fmt.Println(rp.area(), rp.preim()) // Output: 50 30

	// **Key Notes:**
	// - Use a pointer receiver (`*rect`) when:
	//   1. The method needs to modify the original struct.
	//   2. The struct is large, and copying it for each method call is inefficient.
	// - Use a value receiver (`rect`) when:
	//   1. The method does not need to modify the original struct.
	//   2. The struct is small, and copying it is not costly.
}
