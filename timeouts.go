package main

import (
	"fmt"
	"time"
)

// Timeouts ...
func Timeouts() {
	c1 := make(chan string, 1)

	//For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s.
	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "result 1"
	}()

	// Here's the select implementing a timeout.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive fromn c2 will succeed and we will print the result
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// Output:
// timeout 1
// result 2
