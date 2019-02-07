// Closures
package main

import "fmt"

func makeInc(initial, inc int) func() int {
	// Anonymous function: defined inline without name
	return func() int {
		initial += inc
		return initial
	}
}

func main() {
	inc5 := makeInc(0, 5)
	inc3 := makeInc(100, 3)

	fmt.Println("first inc5: ", inc5())
	fmt.Println("first inc3: ", inc3())
	fmt.Println("second inc5:", inc5())
	fmt.Println("second inc3:", inc3())
}
