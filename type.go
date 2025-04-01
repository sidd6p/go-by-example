package main

import "fmt"

// 1. Creating a custom type
type Age int

// 2. Type Aliasing
type MyInt = int

// 3. Defining a struct
type Person struct {
	Name string
	Age  Age
}

// 4. Declaring an interface
type Speaker interface {
	Speak() string
}

// 5. Implementing the interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 6. Creating a function type alias
type MathOperation func(int, int) int

func add(a, b int) int {
	return a + b
}

// 7. Defining constants with a custom type
type Status string

const (
	Active   Status = "Active"
	Inactive Status = "Inactive"
)

func main() {
	// Using custom type
	var myAge Age = 25
	fmt.Println("My Age:", myAge)

	// Using type alias
	var num MyInt = 100
	fmt.Println("MyInt value:", num)

	// Using struct
	p := Person{Name: "Siddhartha", Age: myAge}
	fmt.Println("Person:", p)

	// Using interface
	var s Speaker = Dog{}
	fmt.Println("Dog says:", s.Speak())

	// Using function type alias
	var operation MathOperation = add
	fmt.Println("Sum:", operation(10, 5))

	// Using constants with custom type
	var currentStatus Status = Active
	fmt.Println("Current Status:", currentStatus)
}
