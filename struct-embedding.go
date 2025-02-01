package main

import "fmt"

// Struct embedding in Go allows one struct to include another,
// enabling the outer struct to "inherit" fields and methods from the embedded struct.

type Person struct {
	Name string
	age  int
}

type Employee struct {
	Person
	EmployeeID int
}

// Greet is a method defined on the Person struct.
func (p Person) Greet() {
	fmt.Println("Hello:", p.Name)
}

func main() {

	// Create an instance of Employee, which embeds the Person struct.
	e := Employee{
		Person:     Person{Name: "Siddhartha", age: 10},
		EmployeeID: 10,
	}

	fmt.Println(e)

	// Field Promotion: Fields of the embedded struct (Person) are accessible
	// directly from the outer struct (Employee) without explicitly referencing the embedded struct.
	fmt.Println(e.Name)

	// Method Promotion: Methods of the embedded struct (Person) are promoted
	// and can be called directly from the outer struct (Employee).
	e.Greet()

	// Overriding Methods: The outer struct can define a method with the same name
	// as one in the embedded struct, effectively overriding it.

	// Shadowing Fields: If the outer struct defines a field with the same name
	// as a field in the embedded struct, the outer struct's field takes precedence.
}
