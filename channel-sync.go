package main

import (
	"fmt"
	"time"
)

// ChannelSync Example in golang.
func ChannelSync() {
	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)

	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	msg := <-done
	fmt.Println("Done:", msg)
}

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Done.")

	// Send a value to notify that we’re done.
	done <- true
}
