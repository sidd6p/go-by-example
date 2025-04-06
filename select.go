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

		// The default case makes this select non-blocking.
		// If neither ch1 nor ch2 is ready, this runs immediately.
		default:
			fmt.Println("no message available right now")
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

	// 📝 Explanation of default:
	// The 'default' case in a select makes it non-blocking.
	// That means: if none of the channel operations are ready to proceed,
	// the select will not wait — it will immediately run the default block.
	// This is useful when you want to avoid being stuck waiting on channels.
}
