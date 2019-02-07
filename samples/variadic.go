// Variadic
package main

import "fmt"

// Use arbitrary number of parameters like `fmt.Printf()`
func sum(vals ...int) {
	s := 0
	for _, val := range vals {
		s += val
	}
	fmt.Printf("%v : %d\n", vals, s)
}

func main() {
	sum(1, 2)

	vals := []int{1, 2, 3}
	sum(vals...)

	vals = append(vals, vals...)
	fmt.Println(vals)
}
