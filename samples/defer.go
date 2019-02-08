// Defer
package main

import "fmt"

func deferExample(once bool) {
	defer fmt.Println("first deferred")
	if once {
		fmt.Println("exiting early")
		return
	}
	defer fmt.Println("second deferred")

	fmt.Println("at end of deferExample()")
}

func main() {
	deferExample(false)
	fmt.Println("----------------")
	deferExample(true)
}
