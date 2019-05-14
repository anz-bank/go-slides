package main

import (
	"fmt"
)

// An example that illustrates the basic error creation and the
// fact that error values resulting from errors.New / fmt.Errorf
func main() {
	err := checker(5, 2)
	fmt.Printf("checker(%d, %d): %v\n", 5, 2, err)
	err2 := checker(5, 2)
	fmt.Printf("checker(%d, %d): %v\n", 5, 2, err2)
	fmt.Printf("Comparing error values err == err2: %v\n", err == err2)
}

func checker(m, n int) error {
	if m%n == 0 {
		return nil
	}
	return fmt.Errorf("Input error: m must be divisible by n")
}
