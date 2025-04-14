package main

import (
	"fmt"
	"os"
)

func main() {
	// defer is used to delay execution until the function returns.
	// But it won't run if we use os.Exit before the function ends.
	defer fmt.Println("This will NOT be printed!")

	// Exit the program immediately with status code 3.
	// Any deferred calls like the one above will be skipped.
	os.Exit(3)

	// This line will also not run.
	fmt.Println("This will also NOT be printed.")
}
