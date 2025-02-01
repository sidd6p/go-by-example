package main

import "fmt"

/*
Package-level variables: accessible in all functions unless shadowed by locals.
*/

/*
Key Difference: Constants vs. Variables

1. Constants:
   - Constants in Go do not have a type until used in an expression that requires a type.
   - Their type is inferred based on the context where they are used.
   - Constants cannot be modified after they are declared.
   - Constants can only be inferred to compatible types (e.g., numeric types). Incompatible types like int and string will cause an error.

2. Variables:
   - Variables must be explicitly typed or inferred at the time of declaration.
   - Once declared, the type of a variable is fixed for the remainder of its scope.
   - Variables can be reassigned new values (of the same type).
*/

func main() {
	// Declare a constant (constants do not have a type until used)
	const n = 10 // 'n' is typeless until used

	// Print the constant value
	fmt.Println("Constant n (untyped):", n)

	// Declare variables with explicit and inferred types
	var a = "string"    // Type is inferred as string
	var b, c int = 1, 2 // Declare and initialize multiple variables of the same type
	var d = true        // Type is inferred as bool
	var e int           // Variable declared without initialization; default value is 0

	// Short variable declaration and initialization using `:=` syntax
	f := "apple" // Equivalent to `var f = "apple"`

	// Print all variables with context
	fmt.Println("Variable a (string):", a)
	fmt.Println("Variables b and c (int):", b, c)
	fmt.Println("Variable d (bool):", d)
	fmt.Println("Variable e (int, default value):", e)
	fmt.Println("Variable f (string, declared with :=):", f)

	// Demonstrating constant type inference and compatibility

	// n inferred as int (numeric type), can be assigned to complex128 (compatible numeric type)
	var complexVar complex128 = n
	fmt.Println("n as complex128:", complexVar)

	// n inferred as int (numeric type), cannot be assigned directly to string (incompatible types)
	// Uncommenting the following line would cause an error:
	// var strVar string = n // Error: cannot use n (type int) as type string

	// Example of type conversion: int to float64
	var floatVar float64 = float64(n) // Explicit conversion
	fmt.Println("n as float64:", floatVar)
}
