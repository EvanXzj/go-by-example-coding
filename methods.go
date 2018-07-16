package main

import "fmt"

type rectangle struct {
	width, height float64
}

// This area method has a receiver of *rect.
func (r *rectangle) area() float64 {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

// Methods is a function to illustrate methods in go programming language
func Methods() {
	r := rectangle{10, 6}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
