package main

import "fmt"

/*
i := 1: Declares an integer variable i and initializes it to 1.
&i: Takes the address of i, which is a reference to i.
*iptr: Dereferences iptr to access or modify the value at the address iptr is pointing to.
*/

// zeroVal demonstrates passing by value
func zeroVal(ival int) {
	// Here, 'ival' is a copy of the original variable passed into the function.
	fmt.Println(ival, &ival)
	// Modifying 'ival' does not affect the original variable because it's a copy.
	ival = 0
}

// zeroPtr demonstrates passing by reference using a pointer
func zeroPtr(iptr *int) {
	// 'iptr' is a pointer to an integer, so we dereference it using '*' to get the value it points to.
	// Printing the value the pointer is pointing to, and the address of the pointer.
	fmt.Println(*iptr, iptr)
	// Dereferencing the pointer and changing the value it points to.
	*iptr = 0
}

func getPtr(ival int) *int {
	// Returning the address of 'ival' using the '&' operator in local context.
	return &ival
}

func main() {
	i := 1
	fmt.Println(i, &i, *(&i))

	zeroVal(i)
	// 'i' remains unchanged because zeroVal only modifies the copy of 'i'.
	fmt.Println(i)

	zeroPtr(&i)
	// After calling zeroPtr, 'i' is changed to 0 because the function modifies the value at the address.
	fmt.Println(i)

	println(getPtr(i))
	fmt.Println(&i)
}
