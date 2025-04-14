package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to receive OS signals (like Ctrl+C)
	sigs := make(chan os.Signal, 1)

	// Tell Go to notify this channel when it receives SIGINT (Ctrl+C) or SIGTERM.
	// We're explicitly telling Go *which* signals we want to handle.
	//
	// - SIGINT  is typically sent when a user presses Ctrl+C in the terminal.
	// - SIGTERM is a termination request signal, often used by the OS, Docker, or system tools to
	//   gracefully stop a running program.
	//
	// Note: Go doesn't catch all signals by default. You have to list the ones you want.
	//       Some signals like SIGKILL or SIGSTOP cannot be caught or handled at all.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create another channel to tell the program when it’s safe to exit
	done := make(chan bool, 1)

	// Start a separate goroutine (background task) to wait for a signal
	go func() {
		// Wait for a signal to come in
		sig := <-sigs
		fmt.Println()               // Print a newline for cleaner output
		fmt.Println("Signal:", sig) // Show which signal was received
		done <- true                // Send a value to 'done' to notify main() to continue
	}()

	// Print that we’re waiting for a signal
	fmt.Println("Program is running. Press Ctrl+C to stop.")

	// Wait here (block) until we receive a value from the 'done' channel.
	// We're not interested in the actual value (true/false) — we just use this
	// as a simple way to pause the program until the goroutine says "we're done".
	<-done

	// After receiving the signal, print exit message
	fmt.Println("Program is shutting down.")
}
