package main

import (
	"fmt"
	"sync"
)

func main() {
	// Shared data that multiple goroutines will access
	var counter int

	// Mutex to safely access the shared data
	// A Mutex ensures that only one goroutine can access a shared resource at a time
	var mu sync.Mutex

	// WaitGroup to wait for all goroutines to finish
	// A WaitGroup helps us wait for all the goroutines to complete before continuing
	var wg sync.WaitGroup

	// Number of goroutines to spawn
	numGoroutines := 5

	// Loop to spawn multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment WaitGroup counter (indicating a new goroutine)

		// Launch a new Goroutine
		go func(i int) {
			defer wg.Done() // Decrement WaitGroup counter when this goroutine finishes

			// Lock the mutex to safely update the counter
			// This prevents other goroutines from modifying the counter simultaneously
			mu.Lock()

			// Critical section: updating the shared counter
			// Only one goroutine can execute this at a time due to the mutex lock
			counter++
			fmt.Printf("Goroutine %d updated the counter. Current value: %d\n", i, counter)

			// Unlock the mutex to allow other goroutines to access the shared data
			// After unlocking, other goroutines can enter their critical section
			mu.Unlock()

		}(i) // Passing the loop variable to the goroutine
	}

	// Wait for all goroutines to finish before continuing
	// This ensures the main function waits until all goroutines have completed their work
	wg.Wait()

	// Final counter value after all goroutines have finished
	// The counter will be correctly updated due to mutex protection
	fmt.Println("Final counter value:", counter)
}
