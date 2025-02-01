package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Use reflect to get the type of the variable 'i'
	var i int
	fmt.Println(reflect.TypeOf(i))

	// fmt.Println(i.(type)) // Error: type assertion syntax can't be used here. It's valid only in a type switch.

	switch i := 2; i + 2 {
	case 4, 6: // Check if the result is 4 or 6
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("NULL")
	}

	// Type switch: compares types of values stored in an interface variable (j)
	var j interface{} = 1 // j is an empty interface, which can hold any type (here, it holds an int)

	// Type assertion within a type switch to check the type of j
	switch j.(type) { // Using j.(type) in a type switch to check the actual type stored in j
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	default:
		fmt.Println("No idea")
	}

	// Example to show that type switches only work with interfaces
	// var k int = 10
	// switch k.(type) { // Error: type switches only work with interface{} variables, not direct types like int
	// 	case int:
	// 		fmt.Println("Type is int")
	// }
}
