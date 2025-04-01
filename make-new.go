package main

import "fmt"

/*
   Difference between new() and make():

   1. **new()**:
      - Allocates memory for a variable of a given type and returns a pointer to that type.
      - The variable is initialized to the zero value of the type (e.g., 0 for integers, "" for strings, nil for pointers).
      - It works with all types (including structs, arrays, and basic types).

   2. **make()**:
      - Initializes and returns a reference to a **slice**, **map**, or **channel**.
      - It is used to initialize the internal data structure with specific properties like length, capacity, or buffer size.
      - It does not return a pointer but the actual reference to the data structure itself.

   **In summary**:
   - Use `new()` when you need to allocate memory and get a pointer to a value initialized to its zero value.
   - Use `make()` for slices, maps, and channels to initialize them with specific internal properties (e.g., length, capacity).

*/

func main() {
	// Using new() to allocate memory for a basic type (int)
	// new() returns a pointer to the type and initializes the value to zero (0 for int).
	p := new(int)
	fmt.Println("Using new():")
	fmt.Println("Pointer to int:", p)    // Prints the pointer address
	fmt.Println("Value at pointer:", *p) // Dereference the pointer to get the value (zero value for int is 0)

	// Modify the value through the pointer
	*p = 42
	fmt.Println("Modified value at pointer:", *p) // Prints 42

	// Using new() to allocate memory for a struct (Person)
	type Person struct {
		Name string
		Age  int
	}
	p2 := new(Person) // p2 is a pointer to a Person struct
	fmt.Println("\nUsing new() with structs:")
	fmt.Println("Pointer to Person:", p2)             // Prints the pointer address
	fmt.Println("Zero values of Person struct:", *p2) // Prints the zero values for struct fields (empty string and 0)

	// Modify the fields through the pointer
	p2.Name = "Alice"
	p2.Age = 25
	fmt.Println("Modified Person struct:", *p2) // Prints {Alice 25}

	// Using make() to create and initialize a slice of integers
	// make() initializes the slice with a length of 5 and a capacity of 10
	s := make([]int, 5, 10)
	fmt.Println("\nUsing make() with slice:")
	fmt.Println("Slice:", s)                  // Prints the slice, initially [0, 0, 0, 0, 0]
	fmt.Println("Length of slice:", len(s))   // 5
	fmt.Println("Capacity of slice:", cap(s)) // 10

	// Modify the slice
	s[0] = 100
	fmt.Println("Modified slice:", s) // Prints [100, 0, 0, 0, 0]

	// Using make() to create and initialize a map
	m := make(map[string]int)
	m["age"] = 30
	fmt.Println("\nUsing make() with map:")
	fmt.Println("Map:", m) // Prints map[age:30]

	// Add another key-value pair
	m["year"] = 2025
	fmt.Println("Modified map:", m) // Prints map[age:30 year:2025]

	// Using make() to create and initialize a channel with buffer size of 2
	ch := make(chan int, 2)
	fmt.Println("\nUsing make() with channel:")
	ch <- 10                    // Send a value to the channel
	ch <- 20                    // Send another value to the channel
	fmt.Println("Channel:", ch) // Prints the channel object
	fmt.Println("<-ch:", <-ch)  // Receives 10
	fmt.Println("<-ch:", <-ch)  // Receives 20
}
