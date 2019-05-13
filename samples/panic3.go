package main

import (
	"errors"
	"fmt"
)

func main() {
	oddPanicTest(2)
	oddPanicTest(3)
}

func oddPanicTest(n int) {

	fmt.Printf("odd_panic_test: input: %d, outcome: ", n)
	if err := oddPanic(n); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}

func oddPanic(n int) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()

	if n%2 != 0 {
		panic(fmt.Sprintf("invalid input: %d", n))
	}
	return
}

