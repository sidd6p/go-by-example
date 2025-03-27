package main

import "fmt"

func main() {
	// First loop: while loop equivalent (condition checked first)
	i := 1
	for i <= 3 {
		i = i + 1
	}

	// Second loop: for loop with initialization, condition, and increment
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// The loop below is incomplete; range requires a valid data structure like a slice
	// because range in Go only works with iterables like slices, arrays, maps, strings, or channels.
	// for i := range 3 { }

	// Infinite loop: runs forever, must be manually stopped
	for {

	}
}
