// panic_recover
package main

import (
	"fmt"
)

func main() {
	panicTest()
	fmt.Println("Still alive!")
}

func panicTest() int {
	var n = 10

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()
	panic("panicking")
	return n
}
