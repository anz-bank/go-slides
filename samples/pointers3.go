// Pointers
package main

import "fmt"

func main() {
	var a = [3]*int{}
	var b = []int{1, 2, 3}

	// This looks like we are copying the address of each item into a
	for index, item := range b {
		a[index] = &item
	}

	// Instead we have only copied one address into each three spots
	for _, item := range a {
		fmt.Println(*item)
	}

	fmt.Println()

	// Fix is to access b[i] directly and take its address, rather than through a loop variable
	for index := range b {
		a[index] = &b[index]
	}

	// Now the addresses are copied correctly
	for _, item := range a {
		fmt.Println(*item)
	}
}
