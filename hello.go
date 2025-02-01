package main

/*
Every Go program starts with a package declaration.
The main package is special and serves as the entry point for execution.
The main function, defined in the main package, is automatically executed by the Go runtime.
It takes no parameters, returns no values, and cannot be called directly.
*/

import "fmt"

func main() {
	fmt.Println("Hello Earth")
	fmt.Println("go" + "lang")
}
