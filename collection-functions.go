package main

import (
	"fmt"
	"strings"
)

// We often need our programs to perform operations on collections of data,
// like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.

func collectionFunctions() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(index(strs, "pear"))
	fmt.Println(include(strs, "grape"))

	fmt.Println(any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(all(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(mapping(strs, strings.ToUpper))
}

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

func any(vs []string, fn func(string) bool) bool {
	for _, v := range vs {
		if fn(v) {
			return true
		}
	}

	return false
}

func all(vs []string, fn func(string) bool) bool {
	for _, v := range vs {
		if !fn(v) {
			return false
		}
	}

	return true
}

func filter(vs []string, fn func(string) bool) []string {
	vsf := make([]string, 0)

	for _, v := range vs {
		if fn(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func mapping(vs []string, fn func(string) string) []string {
	vsm := make([]string, len(vs))

	for i, v := range vs {
		vsm[i] = fn(v)
	}

	return vsm
}
