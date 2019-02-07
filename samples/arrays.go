// Arrays
package main

import "fmt"

func main() {
	// Arrays are of fixed size
	var a [3]int
	fmt.Println("initial a:", a)

	a[1] = 10
	fmt.Println("set a:", a)
	fmt.Println("a[1]:", a[1])
	fmt.Println("len:", len(a))

	// Short declarations and initialisation
	b := [2]string{"a", "b"}
	fmt.Println("initial b:", b)

	// Two dimensional array
	c := [2][3]int{{1, 2, 3}, {10, 20, 30}}
	fmt.Println("initial c: ", c)
}
