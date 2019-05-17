package main

import (
	"errors"
	"fmt"
)

var NotEqualErr = errors.New("Inputs not equal")

func main() {
	err1 := modcheck(5, 3)
	err2 := modcheck(5, 3)
	fmt.Println("err1 == err2: ", err1 == err2)
	err3 := eqcheck(1, 1)
	err4 := eqcheck(2, 2)
	fmt.Println("err3 == err4: ", err3 == err4)
}

func modcheck(m, n int) error {
	if m%n == 0 {
		return nil
	}
	return fmt.Errorf("Input error: %d mod %d must be 0", m, n)
}

func eqcheck(m, n int) error {
	if m != n {
		return NotEqualErr
	}
	return nil
}
