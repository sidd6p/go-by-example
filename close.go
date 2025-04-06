package main

import "fmt"

// producer is a function that sends integers to the channel and then closes it
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // send value to the channel
	}
	close(ch) // close the channel after sending all values
}

func main() {
	ch := make(chan int) // create a new unbuffered channel of type int

	go producer(ch) // start the producer in a separate goroutine

	// range loop receives values from the channel until it is closed
	for val := range ch {
		fmt.Println(val) // print the received value
	}

	fmt.Println("Channel closed, done receiving.")
}
