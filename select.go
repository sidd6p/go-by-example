package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two unbuffered channels of type string
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Launch a goroutine to send a message on ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	// Launch another goroutine to send a message on ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// Loop runs exactly two times
	// So only two messages will be received in total, regardless of how many are sent
	for i := 0; i < 2; i++ {
		select {
		// If a message is available on ch1, this case is executed
		case msg1 := <-ch1:
			fmt.Println(msg1)

		// If a message is available on ch2, this case is executed
		case msg2 := <-ch2:
			fmt.Println(msg2)

			// If neither channel has data yet, select blocks and waits
			// (unless you add a 'default' case)
		}
	}

	// ⚠️ Note: If more messages are sent on ch1 or ch2, they will not be received here.
	// The program exits after receiving just 2 messages.
	// To keep listening forever, you can use:
	//
	// for {
	//     select {
	//     case msg1 := <-ch1:
	//         fmt.Println(msg1)
	//     case msg2 := <-ch2:
	//         fmt.Println(msg2)
	//     }
	// }
}
