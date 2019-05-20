package main

import "fmt"

// Defer functions are executed in LIFO order
func f1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

// The arguments to the deferred function are evaluated when defer executes
// not when the call executes
func f2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

// defer call executes after evaluating the return but before actually returning
func f3() (result int) {
	defer func() {
		result *= 7
	}()
	return 6
}

func main() {
	fmt.Print("f1: ")
	f1()
	fmt.Println()
	fmt.Print("f2: ")
	f2()
	fmt.Println()
	fmt.Print("f3: ", f3())
	fmt.Println()
}
