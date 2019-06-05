// error_compare
package main

import (
	"errors"
	"fmt"
)

var ErrNotEqual = errors.New("not equal")

func main() {
	fmt.Printf("equal(2, 3) == equal(2, 3): %v", equal(2, 3) == equal(2, 3))
}

func equal(m, n int) error {
	if m != n {
		// return fmt.Errorf("%d and %d are not equal", m, n)
		return ErrNotEqual
	}
	return nil
}
