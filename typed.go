package main

import "fmt"

func main() {
	// ✅ Explicitly typed variables:
	var a int = 10          // 'a' is explicitly typed as int
	var b float64 = 3.14    // 'b' is explicitly typed as float64
	var c string = "GoLang" // 'c' is explicitly typed as string

	// ✅ Type-inferred (still typed, just not explicitly written):
	d := 20          // 'd' is inferred to be of type int
	e := 2.71        // 'e' is inferred to be float64
	f := "Hello Go!" // 'f' is inferred to be string

	// ✅ Untyped constants:
	const pi = 3.1415 // 'pi' is untyped until it's used
	const answer = 42 // 'answer' is also untyped until used

	// When we assign untyped constants to typed variables,
	// Go automatically converts them to the target type:
	var precisePi float64 = pi      // 'pi' becomes float64 here
	var ultimateAnswer int = answer // 'answer' becomes int here

	// ✅ Typed channel (channels are always explicitly typed)
	ch := make(chan int) // 'ch' is a channel that only carries int values

	// Display all the values
	fmt.Println("Typed variables:")
	fmt.Println("a (int):", a)
	fmt.Println("b (float64):", b)
	fmt.Println("c (string):", c)

	fmt.Println("\nInferred types:")
	fmt.Println("d (int):", d)
	fmt.Println("e (float64):", e)
	fmt.Println("f (string):", f)

	fmt.Println("\nUntyped constants assigned to typed vars:")
	fmt.Println("precisePi (float64):", precisePi)
	fmt.Println("ultimateAnswer (int):", ultimateAnswer)

	fmt.Println("\nTyped channel created: ch := make(chan int)")
}

/*
💡 Summary:

1. Go is a statically typed language — every variable has a type at compile time.
2. You can declare variables with explicit types (e.g., `var a int`) or let Go infer the type (e.g., `a := 10`).
3. Constants in Go are untyped by default and adapt to the type of the variable they are assigned to.
4. Channels, structs, arrays, etc., are always explicitly typed — you define what kind of data they work with.

The goal is to balance strict type safety with clean, readable code.
*/
