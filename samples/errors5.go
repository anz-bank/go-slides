package main

import (
	"errors"
	"fmt"
)

var ErrNotEqual = errors.New("not equal")

func main() {
	err1 := eqcheck(1, 1)
	err2 := eqcheck(2, 3)
	err3 := eqcheck(2, 3)
	fmt.Printf("err1: %v, err2: %v, err2 == err3: %v\n", err1, err2, err2 == err3)
}

func eqcheck(m, n int) error {
	if m != n {
		return ErrNotEqual
	}
	return nil
}
