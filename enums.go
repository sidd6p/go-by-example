package main

import "fmt"

// Define a custom type 'Day' based on the int type.
// Type keyword is used to define Creating Custom Types, Type Aliasing, Defining Structs, Declaring Interfaces, Creating Function Type Aliases or Defining Constants with Custom Types
type Day int // aliasing mechanism

// Define constants for days of the week using iota.
const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Predefine an array of day names as a package-level variable.
// This avoids recreating the array every time the String() method is called.
var dayNames = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

/*
If a type implements the String() method,
it defines how values of that type will be converted to a string when using functions like fmt.Println() or fmt.Sprintf().
*/
// String method defines how values of type 'Day' will be converted to a string.
func (d Day) String() string {
	return dayNames[d] // Use the predeclared array for the conversion.
}

func main() {
	var today Day = Wednesday
	fmt.Println(today) // Output: Wednesday
}
