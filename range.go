package main

import (
	"fmt"
)

// Range iterate over elements in a variety of data stuctures
func Range() {
	// Here we use range to sum the numbers in a slice. Array work like this too.
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for index, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", index)
		}
	}

	// range on map iterates over key/value pairs
	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for key, value := range m {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
