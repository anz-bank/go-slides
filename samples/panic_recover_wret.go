// panic_recover_wret
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("validateOdd(2); err = ", validateOdd(2))
	fmt.Println("validateOdd(3); err = ", validateOdd(3))
}

func validateOdd(n int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()
	oddPanic(n)
	return
}

func oddPanic(n int) {
	if n%2 != 0 {
		panic(fmt.Sprintf("invalid input: %d", n))
	}
	return
}
