package main

import (
	"fmt"
	"time"
)

// f prints a string along with an index three times.
func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	// Using 'go' keyword to invoke f as a goroutine.
	// This runs concurrently with the main function.
	go f("goroutine")

	// Directly calling f. This runs in the main goroutine.
	f("Direct")

	// Anonymous function as a goroutine.
	// The message is passed as a parameter to the anonymous function.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Adding a sleep to allow goroutines enough time to execute
	// before the program terminates. Without this, some goroutines
	// might not finish as the main function exits.
	time.Sleep(time.Second)

	// Prints after all other operations.
	fmt.Println("done")
}
