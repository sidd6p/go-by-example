package main

import (
	"errors"
	"fmt"
)

// Custom error type `myError` with additional fields
type myError struct {
	code int
	msg  string
}

// Implement the `Error` method for `myError` to satisfy the `error` interface
func (e *myError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.msg)
}

// Function that returns an error (either custom or built-in)
func doSomething(fail bool) error {
	if fail {
		return &myError{404, "Resource not found"} // Returning a custom error
	}
	return errors.New("A generic error occurred") // Returning a standard error
}

// Function to demonstrate type assertion on an interface{}
func typeAssertionDemo(value interface{}) {
	// Trying to assert value as a string
	str, ok := value.(string)
	if ok {
		fmt.Println("Extracted string:", str)
	} else {
		fmt.Println("Failed to extract string") // If assertion fails
	}

	// Trying to assert value as an int
	num, ok := value.(int)
	if ok {
		fmt.Println("Extracted integer:", num)
	} else {
		fmt.Println("Failed to extract integer")
	}
}

// Function demonstrating a type switch
func typeSwitchDemo(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an integer:", v)
	case bool:
		fmt.Println("It's a boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// --- Example 1: Basic Type Assertion ---
	fmt.Println("---- Type Assertion Example ----")

	// Declaring an empty interface variable 'x' that can hold any type
	// The empty interface is a type that can store any value
	var x interface{} = "Hello, Go!" // Assigning a string value to the empty interface
	typeAssertionDemo(x)

	// --- Example 2: Type Assertion on Errors ---
	fmt.Println("\n---- Type Assertion on Errors ----")
	err := doSomething(true) // Simulating an error

	if err != nil {
		// Type assertion to check if the error is of type *myError
		if e, ok := err.(*myError); ok {
			fmt.Println("Custom error details:")
			fmt.Println("Code:", e.code) // Extracting specific fields
			fmt.Println("Message:", e.msg)
		} else {
			fmt.Println("Standard error:", err)
		}
	}

	// --- Example 3: Using Type Switch ---
	fmt.Println("\n---- Type Switch Example ----")
	typeSwitchDemo(42)       // Integer case
	typeSwitchDemo("Hello!") // String case
	typeSwitchDemo(true)     // Boolean case
	typeSwitchDemo(3.14)     // Default case (unknown type)
}
