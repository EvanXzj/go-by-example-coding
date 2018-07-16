package main

import "fmt"

// Recursive illustrate a recursive example
func Recursive() {
	fmt.Println(factorial(7))
}

// Go supports recursive functions. Here’s a classic factorial example.
func factorial(n int) int {
	// This fact function calls itself until it reaches the base case of fact(0).
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
