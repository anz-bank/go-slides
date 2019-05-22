// panic_recover.go
package main

import (
	"fmt"
)

func main() {
	n := panicTest()
	fmt.Printf("panicTest: %d\n", n)
}

func panicTest() int {
	var n int = 10

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("failed")
	return n
}
