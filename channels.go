package main

import (
	"fmt"
	"time"
)

func main() {
	// 🔹 Unbuffered Channel Example
	unbuffered()

	fmt.Println()

	// 🔹 Buffered Channel Example
	buffered()
}

// -----------------------------------
// 🔄 Unbuffered Channel Example
func unbuffered() {
	// Create a new unbuffered channel of type string.
	messages := make(chan string) // Unbuffered channel (capacity = 0)

	// Start a new goroutine to send a message.
	go func() {
		messages <- "ping" // This send will block until another goroutine is ready to receive.
	}()

	// Receive a value from the channel.
	msg := <-messages // Blocks until the value is sent.

	fmt.Println("Unbuffered:", msg)
}

// -----------------------------------
// 📦 Buffered Channel Example
func buffered() {
	// Create a buffered channel with capacity of 2.
	messages := make(chan string, 2)

	// Send two values to the buffered channel (no blocking yet).
	messages <- "hello"
	fmt.Println("Buffered: sent 'hello'")
	messages <- "world"
	fmt.Println("Buffered: sent 'world'")

	// This third send will block because buffer is full (capacity = 2).
	go func() {
		fmt.Println("Buffered: trying to send 'Go' (will block until space is available)...")
		messages <- "Go"
		fmt.Println("Buffered: sent 'Go' after buffer had space")
	}()

	time.Sleep(2 * time.Second)

	// Read one value to free up space in the buffer.
	fmt.Println("Buffered: received ->", <-messages)

	time.Sleep(1 * time.Second)
	fmt.Println("Buffered: received ->", <-messages)
	fmt.Println("Buffered: received ->", <-messages)
}

/*
📝 Explanation:

🔄 Unbuffered Channel:
- Created using `make(chan string)`
- Has no internal capacity → send blocks until receive is ready.
- Ensures synchronization between goroutines.
- Acts like a handshake.

📦 Buffered Channel:
- Created using `make(chan string, N)` where N is the buffer size.
- Can hold up to N values before blocking the sender.
- Sending blocks ONLY when the buffer is full.
- Receiving blocks ONLY when the buffer is empty.

Important:
- If the buffer is full and you try to send more, the sender goroutine will BLOCK until space is freed.
- For unbuffered channels, you can think of capacity as **0** (not 1).
- Sending on an unbuffered channel blocks until a receiver is ready.
*/
