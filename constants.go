package main

import (
	"fmt"
	"math"
)

// const declares a constant value
const str string = "constant"

// Constant details constants in golang
func Constant() {
	fmt.Println(str)

	//  A const statement can appear anywhere a var statement can
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric const has no type until it's given one, such as by an explicit cast
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
