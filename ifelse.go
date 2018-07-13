package main

import (
	"fmt"
)

// ElseIf details else if statement in golang
func ElseIf() {
	// Basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditional;
	// any variables declared in this statement are aviable in all branches/
	if num := 0; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
	// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
	// ternary: adj. [数] 三元的，三重的；[数] 三进制的；第三的
}
