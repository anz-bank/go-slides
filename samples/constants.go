// Constants
package main

import "fmt"

const Phi = 2.61803398875

// The `iota` identifier is used in const declarations to
// simplify definitions of incrementing numbers.
const (
	c0 = iota // c0 == 0
	c1        // c1 == 1
	c2        // c2 == 2
)

func main() {
	const World = "ವಿಶ್ವ"
	fmt.Println("Hello", World)
	fmt.Println("Beauty in number:", Phi)

	fmt.Println(c0, c1, c2)
}
