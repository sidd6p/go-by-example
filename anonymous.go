package main

import "fmt"

func main() {
	// 1. Declaring and immediately invoking an anonymous function
	func() {
		fmt.Println("Hello from anonymous function!")
	}() // Immediately invoking it

	// 2. Storing an anonymous function in a variable and calling it
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Siddhartha") // Calling the function via the variable

	// 3. Anonymous function with a return value
	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 4) // Calling the function
	fmt.Println("Sum:", result)

	// 4. Using an anonymous function as a parameter (Closure)
	apply := func(fn func(int, int) int) {
		result := fn(5, 3)
		fmt.Println("Result from function:", result)
	}
	apply(func(a, b int) int { return a * b }) // Passing anonymous function to apply

	// 5. Declaring an anonymous function and defining it later
	var greetLater func(string)

	// Defining the anonymous function later
	greetLater = func(name string) {
		fmt.Printf("Hello again, %s!\n", name)
	}
	greetLater("Siddhartha") // Calling the function later
}
