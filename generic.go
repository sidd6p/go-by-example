package main

import "fmt"

func main() {
	// Using the generic function
	intResult := Add(10, 20)     // T is int
	floatResult := Add(5.5, 2.3) // T is float64
	fmt.Println("Sum of integers:", intResult)
	fmt.Println("Sum of floats:", floatResult)

	// Using the generic type
	stringPair := Pair[string, string]{First: "Hello", Second: "World"}
	intPair := Pair[int, int]{First: 1, Second: 2}

	// Using 'any' as a type constraint for Pair
	mixedPair := Pair[int, any]{First: 1, Second: "Go"}

	fmt.Println("String Pair:", stringPair)
	fmt.Println("Integer Pair:", intPair)
	fmt.Println("Mixed Pair:", mixedPair)
}

// 1. Type Constraints: Define a custom constraint interface
type Number interface {
	int | float64
}

// 2. Generic Type: Define a generic struct
type Pair[T any, U any] struct {
	First  T
	Second U
}

// Generic Function: Uses Type Parameters and Type Constraints
func Add[T Number](a, b T) T {
	return a + b
}
