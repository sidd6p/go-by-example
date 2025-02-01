package main

import "fmt"

const (
	A = iota // iota starts at 0 in this const block
	B        // increments to 1 for the next constant
	_        // increments to 2, but the value is skipped by assigning it to _
	C        // increments to 3
)

const (
	a = iota * 10 // iota starts at 0; 0 * 10 = 0
	b             // increments to 1; 1 * 10 = 10
)

func main() {
	fmt.Println(A, B, C) // Output: 0 1 3

	fmt.Println(a, b) // Output: 0 10
}
