// Slices
package main

import "fmt"

func main() {

	array := [6]int{1, 1, 2, 3, 5, 8}

	// A slice is a dynamically-sized, flexible view into the elements of an array.
	var slice1 []int = array[1:4]
	var slice2 []int = array[:2]
	var slice3 []int = array[3:]
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3, "\n")

	// Slices can be created with the built-in `make` function.
	a := make([]int, 3)
	b := make([]int, 3, 10)
	c := []int{1, 2, 3}
	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))
	fmt.Printf("c: %v, len: %d, cap: %d\n\n", c, len(c), cap(c))

	// Append
	s := []string{"C", "D", "E"}
	s = append(s, "f")
	s = append(s, "g", "h", "i")
	s = append([]string{"a", " b"}, s...)
	fmt.Printf("s: %v, len: %d, cap: %d\n\n", s, len(s), cap(s))

	// Copy and reference
	cs1 := make([]string, len(s))
	copy(cs1, s)
	rs1 := s
	rs2 := s[:]
	s[0] = "Z"
	fmt.Println("cs1:", cs1)
	fmt.Println("rs1:", rs1)
	fmt.Println("rs2:", rs2)
	fmt.Println("s  :", s)
}
