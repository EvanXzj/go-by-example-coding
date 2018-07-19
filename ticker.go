package main

import (
	"fmt"
	"time"
)

// Ticker func illuste the use of tickers in golang
func Ticker() {
	// Timers are for when you want to do something once in the future -
	// tickers are for when you want to do something repeatedly at regular intervals.
	// Here’s an example of a ticker that ticks periodically until we stop it.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Ticker can be stopped like timers
	// Once a ticker is stopped it won't receive any more values on its channel. we will stop it after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	// When we run this program the ticker should tick 3 times before we stop it.
}
