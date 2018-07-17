package main

import (
	"fmt"
)

/*
*	Channels are the pipes that connect concurrent goroutines
*	You can send values into channels from one goroutines and recive those vlaues into annother goroutines
* */

// Channels illustrate channels in go
func Channels() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above.
	go func() {
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel
	// Here we'll recevie the  "ping" message we send above and print it out
	msg := <-messages
	fmt.Println(msg)
}
