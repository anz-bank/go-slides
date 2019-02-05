// Variables
package main

import "fmt"

var (
	idx   int     // global variable
	two   int = 2 // with initialiser
	three     = 3 // with inferred type
)

func main() {
	fmt.Printf("idx  - Val: %v, Type: %T\n", idx, idx)
	fmt.Printf("two  - Val: %v, Type: %T\n", two, two)
	fmt.Printf("three- Val: %v, Type: %T\n\n", three, three)

	var i int
	var yes bool = true
	var nope = false
	fmt.Printf("i    - Val: %v, Type: %T\n", i, i)
	fmt.Printf("yes  - Val: %v, Type: %T\n", yes, yes)
	fmt.Printf("nope - Val: %v, Type: %T\n\n", nope, nope)

	// Short variable declarations, only inside functions
	s := "shanti"
	fmt.Printf("s    - Val: %v, Type: %T\n", s, s)
}
