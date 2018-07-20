package main

// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn’t occur during normal operation,
// or that we aren’t prepared to handle gracefully.

// Panic details painc in golang
func Panic() {
	panic("a problem")

	// fmt.Println("painc")  // unreachable code
}
