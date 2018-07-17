package main

import (
	"fmt"
	"time"
)

// GoRoutines function illustrate go routines in golang
func GoRoutines() {
	// A goroutine is a lightweight thread of execution.
	// Suppose we have function call fn(s). Here's how we'd call that in usual way, running it synchronously.
	fn("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go fn("goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now,
	// so execution falls through to here.
	time.Sleep(2 * time.Second)
	fmt.Println("Done")

	// When we run this program, we see the output of the blocking call first,
	// then the interleaved output of the two goroutines(run it multiple times).
	// This interleaving reflects the goroutines being run concurrently by the Go runtime.
}

func fn(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
	}
}
