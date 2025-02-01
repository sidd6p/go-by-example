package main

import "fmt"

// getInt is a function that returns another function (a closure).
func getInt() func() int {
	i := 0
	// Returning an anonymous function (closure) that modifies 'i' when called
	return func() int {
		i += 1 // This function accesses 'i' from the enclosing scope and increments it
		return i
	}
}

func main() {
	// 'v1' is a closure returned by calling getInt().
	v1 := getInt()

	// Calling the closure multiple times will increment the value of 'i' each time.
	fmt.Println(v1()) // Outputs: 1
	fmt.Println(v1()) // Outputs: 2
	fmt.Println(v1()) // Outputs: 3

	// It has its own independent 'i', which starts from 0.
	v2 := getInt()
	fmt.Println(v2()) // Outputs: 1 (independent of v1's closure)

	// An anonymous recursive function to calculate the Fibonacci number
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2) // Since fib was previously declared in main, Go knows which function to call with fib here.
	}

	fmt.Println(fib(7))
}
