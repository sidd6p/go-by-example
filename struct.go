package main

import "fmt"

type person struct {
	name string
	age  int
}

// `newPerson` demonstrates returning a pointer to a local variable.
// The local variable `p` is safe to use even after the function exits
// because Go's garbage collector ensures it stays in memory as long as there is a reference to it.
func newPerson(name string) *person {
	// Declare a local variable `p` of type `person`.
	p := person{name: name}
	// Assign a default value to the `age` field.
	p.age = 42
	// Return a pointer to the local variable `p`.
	// Go ensures that `p` is moved to the heap if needed to keep it accessible.
	return &p
}

func main() {
	// This is a value type. A copy of the struct is created.
	fmt.Println(person{name: "Alice", age: 30}) // Prints: {Alice 30}

	// The `age` field will be initialized to its zero value (0 for int).
	fmt.Println(person{name: "Fred"}) // Prints: {Fred 0}

	// Create a `person` instance and return its memory address (pointer).
	// The `&` operator creates a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40}) // Prints: &{Ann 40}

	// Create a `person` instance using the `newPerson` function.
	// The `newPerson` function returns a pointer to the struct.
	fmt.Println(newPerson("Jon")) // Prints: &{Jon 42}

	// Access fields of a `person` instance (value type).
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Access the `name` field. Prints: Sean

	// Use a pointer to access and modify fields of a `person` instance.
	sp := &s            // Pointer to the `s` instance
	fmt.Println(sp.age) // Access the `age` field via pointer. Prints: 50
	sp.age = 51         // Modify the `age` field via pointer.
	fmt.Println(sp.age) // Prints the updated `age`. Prints: 51

	// Define and initialize an anonymous struct for a `dog`.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex", // `name` field.
		true,  // `isGood` field.
	}
	fmt.Println(dog) // Print the anonymous struct instance. Prints: {Rex true}

	// Additional example to demonstrate value vs. pointer differences.
	fmt.Println("\n--- Value vs Pointer Example ---")

	// 	This creates a value of the person struct.
	// A copy of the person struct is passed to fmt.Println.
	// Struct as value (no pointer).
	p1 := person{name: "Ann", age: 40}
	fmt.Println(p1) // Prints: {Ann 40}

	// 	This creates a pointer to the person struct using the & operator.
	// The memory address of the person instance is passed to fmt.Println.
	// Struct as pointer.
	p2 := &person{name: "Ann", age: 40}
	fmt.Println(p2) // Prints: &{Ann 40}

	// Modifying the struct through the pointer.
	p2.age = 41     // Works because p2 is a pointer.
	fmt.Println(p2) // Prints: &{Ann 41}

	// Modifying the value type requires direct reassignment.
	p1.age = 42     // Modifies the value directly.
	fmt.Println(p1) // Prints: {Ann 42}
}
