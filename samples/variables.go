// Variables
package main

import "fmt"

var j int // single variable declaration

var ( // declaring multiple variables
	rat int     // global variable
	dog int = 2 // with initialiser
	foo     = 3 // with inferred type
)

func main() {
	fmt.Printf("j    - Val: %v, Type: %T\n\n", j, j)
	fmt.Printf("rat  - Val: %v, Type: %T\n", rat, rat)
	fmt.Printf("dog  - Val: %v, Type: %T\n", dog, dog)
	fmt.Printf("foo  - Val: %v, Type: %T\n\n", foo, foo)

	var i int
	var yes bool = true
	var nope = false
	fmt.Printf("i    - Val: %v, Type: %T\n", i, i)
	fmt.Printf("yes  - Val: %v, Type: %T\n", yes, yes)
	fmt.Printf("nope - Val: %v, Type: %T\n\n", nope, nope)

	// short variable declarations, only inside functions
	s := "shanti"
	fmt.Printf("s    - Val: %v, Type: %T\n\n", s, s)

	// single line declaration
	a, b, c := 1, 2, 3
	fmt.Printf("a    - Val: %v, Type: %T\n", a, a)
	fmt.Printf("b    - Val: %v, Type: %T\n", b, b)
	fmt.Printf("c    - Val: %v, Type: %T\n", c, c)
}
