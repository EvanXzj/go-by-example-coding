package main

import "fmt"

// Closure describe anonymous functions in golang
func Closure() {
	// We call intSeq func, assigning the result (a function) to nextInt variable
	// This function value captures its own i value, which will be updated each time we call nextInt
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++

		return i
	}
}
