package main

import (
	"fmt"
)

// Variables is a declaration example
func Variables() {
	// declare one or more variables
	var name = "Chui Dylan"
	fmt.Println(name)

	// declare multiple variables at once
	var a, b int = 0, 1
	fmt.Println(a, b)

	// Go will infer the type of initialized variables
	var c = true
	fmt.Println(c)

	// variables declared without a corresponding initialization are Zero-Valued.
	// int -> 0
	// string -> ""
	// boolean -> false
	var d int
	var e bool
	fmt.Println(d, e)

	// The syntax is shorthand for declaring and initializing a variable
	// In this case var f string = "short"
	f := "short"
	fmt.Println(f)
}
