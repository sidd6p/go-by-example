// Go program to illustrate the method
// with non-struct type receiver
package main

import "fmt"

// Type definition
type data int

// Defining a method with
// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

/*
The commented code demonstrates that you cannot define methods directly on built-in types (like int). This would result in a compiler error because Go only allows methods with receivers of defined types (such as structs or custom types).
In this case, int is a built-in type, and Go does not allow defining methods for it.
func(d1 int)multiply(d2 int)int{
return d1 * d2
}
*/

// Main function
func main() {
	value1 := data(23)
	value2 := data(20)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}
