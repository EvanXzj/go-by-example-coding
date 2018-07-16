package main

import (
	"fmt"
)

type person struct {
	name string
	male bool
	age  int
}

// Struct Function to illustrate go's structs
func Struct() {
	// This syntax creates a new struct
	fmt.Println(person{"chuidylan", true, 27})

	// You can name the fields when initializing a struct
	fmt.Println(person{age: 27, name: "chuidylan", male: true})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred", age: 10})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with dot syntax
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	// Structs are mutable.
	fmt.Println(sp.age)
}
