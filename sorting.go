package main

import (
	"fmt"
	"sort"
)

// Sorting details sort package in golang
func Sorting() {
	// Sort methods are specific to the builtin type;
	// here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints
	ints := []int{7, 3, 5}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// We can also use sort to check if a slice is already in sorted order.
	unSortInts := []int{12, 1, 5}
	a := sort.IntsAreSorted(ints)
	b := sort.IntsAreSorted(unSortInts)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
