package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

// Slices describe slices type in golang
func Slices() {
	// Unlike arrays, slice are typed only by the elements they contain (not include the the number of elements)
	// To create an empty slice with non-zero length, use the builtin make function.
	// Here we make a slice of strings of length 3 (initially zero-valued)
	s := make([]string, 3)
	fmt.Println("empty slice:", s)

	// We can set an get element value just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// The builtin len method returns the length of the slice
	fmt.Println("len: ", len(s))

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// Slices can also be copy'd
	// Here we create an empty slice c of the same length as s and copy into c form s
	c := make([]string, 3)
	copy(c, s)
	fmt.Println("after copy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high], not inlcude the element at high position.
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1: ", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2: ", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3: ", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("declare a slice in a single line: ", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// To append one slice to another, use ... opeator to expand the second argument to a list of arguments.
	a := []string{"Ismail", "BahaEddine"}
	b := []string{"Pogba", "Mane"}

	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	fmt.Println("append two slices:", a)

	// Some gotcha
	// http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
	fmt.Println("found digits: ", findDigits("./finddigits.txt"))

	// output:
	// found digits: [49 50 51]
	// [49 50 51] -> [1 2 3]
}

func findDigits(path string) []byte {
	// ignore the error
	bytes, _ := ioutil.ReadFile(path)
	bytes = digitRegexp.Find(bytes)

	result := make([]byte, len(bytes))
	copy(result, bytes)

	return result
	// or return bytes
}
