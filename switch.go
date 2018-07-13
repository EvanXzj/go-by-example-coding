package main

import (
	"fmt"
	"time"
)

// Switch function details switch statement in golang
func Switch() {
	// basic switch
	i := 2
	fmt.Print("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Twon")
	case 3:
		fmt.Println("Three")
	}

	// You can use commas to separate multiple expressions in the same case statement
	// We use the optional default case in this example as well
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")

	default:
		fmt.Println("It's a weekday")
	}

	// Switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constans.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value
	/// In this example, the variable will have the tpe corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(0)
	whatAmI(1 + 2i)
	whatAmI("string")
}
