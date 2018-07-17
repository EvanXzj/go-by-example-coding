package main

import (
	"errors"
	"fmt"
)

// Errors function details error in golang
func Errors() {
	// The two loops below test out each of our error-returning functions.
	// Note that the use of an inline error check on the if line is a common idiom in Go code.

	for _, i := range []int{7, 42} {
		if r, e := fn1(i); e != nil {
			fmt.Println("fn1 failed: ", e)
		} else {
			fmt.Println("fn1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := fn2(i); e != nil {
			fmt.Println("fn2 failed: ", e)
		} else {
			fmt.Println("fn2 worked: ", r)
		}
	}

	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type via type assertion.
	_, e := fn2(42)

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	}
}

// By convention, errors are the last return value and have type error, a built-in interface.
func fn1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("fn1 can't work with 42")
	}

	return arg + 3, nil
}

// It’s possible to use custom types as errors by implementing the Error() method on them.
// Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func fn2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "fn2 can't work with 42"}
	}

	return arg + 3, nil
}
