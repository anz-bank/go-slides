package main

import (
	"fmt"
)

func f1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func f2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}

// returns 42
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
}
