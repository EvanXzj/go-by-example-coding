package main

import (
	"fmt"
	"sort"
)

type byLength []string

// Implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// SortingByFunctions illustrate sorting by a custom function in Go
func SortingByFunctions() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))

	fmt.Println(fruits)

	// By following this same pattern of creating a custom type,
	// implementing the three Interface methods on that type, and then calling sort.Sort on a collection of that custom type,
	// we can sort Go slices by arbitrary functions.
}
