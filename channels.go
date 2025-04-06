package main

import "fmt"

func main() {
	// Create a new channel of type string.
	// Channels are typed, so this one will carry string values.
	messages := make(chan string)

	// Start a new goroutine (lightweight thread).
	// Inside it, send the string "ping" into the channel using the send syntax: channel <- value
	go func() {
		messages <- "ping" // This send will block until another goroutine is ready to receive.
	}()

	// Receive a value from the channel using the receive syntax: <-channel
	// Since channels block by default, this line will wait until something is sent to the channel.
	msg := <-messages

	// Print the received message to the console.
	fmt.Println(msg)
}

/*
Explanation:
- Channels are used to communicate between goroutines safely.
- `make(chan string)` creates an unbuffered channel (capacity 0) for string communication.
- The `go` keyword starts a new goroutine, allowing the anonymous function to run concurrently.
- Sending and receiving on an unbuffered channel are both blocking.
- This ensures synchronization: the main goroutine waits for the value to be sent before continuing.
- Output will be: ping
*/
