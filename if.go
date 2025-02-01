package main

/*
There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
*/

func main() {
	// Basic if-else statement (condition check without initialization)
	if 7%4 == 2 {
	} else {
	}

	// If-else with variable initialization inside the if condition
	// num is scoped only within the if block
	if num := 9; num < 0 {
	} else if num < 10 {
	} else {
	}
}
