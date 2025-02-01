package main

import (
	"errors"
	"fmt"
)

// f1 demonstrates simple error handling using the built-in errors.New.
func f1(num float64) (float64, error) {
	if num == 0 {
		// Return an error indicating the issue
		return -1, errors.New("it is zero")
	}
	// Return the reciprocal and no error
	return float64(1.0 / num), nil
}

// Custom error type `myError` allows for more detailed error handling
// with additional fields like `num` and `prob` (problem description).
type myError struct {
	num  int    // The number causing the error
	prob string // A description of the problem
}

// Error method implements the error interface for `myError`,
// providing a formatted error message.
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.num, e.prob)
}

// f2 demonstrates the use of a custom error type `myError`.
// It returns the reciprocal of an integer and an error if the input is zero.
func f2(num int) (int, error) {
	if num == 0 {
		// Return a custom error when the input is zero
		return -1, &myError{num, "it is zero"}
	}
	// Return the reciprocal and no error
	return int(1.0 / float64(num)), nil
}

func main() {
	// Example usage of f1
	fmt.Println(f1(10.0)) // Should print: (0.1, <nil>)
	fmt.Println(f1(0))    // Should print: (-1, it is zero)

	// Example usage of f2
	fmt.Println(f2(10)) // Should print: (0, <nil>)
	fmt.Println(f2(0))  // Should print: (-1, 0 - it is zero)

	// Handling the custom error returned by f2
	_, e := f2(0)
	fmt.Println("This is me", e.Error()) // Accessing the error message
	if ae, ok := e.(*myError); ok {      // Type assertion to access custom error fields
		fmt.Println(ae.num)  // Prints: 0
		fmt.Println(ae.prob) // Prints: it is zero
	}
}
