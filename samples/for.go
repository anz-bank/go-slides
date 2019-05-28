// For
package main

import "fmt"

func main() {
	sum := 0

	// Basic for loop - like C, Java, or JS but without parentheses:
	// init statement, condition expression, post statement
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum 1:", sum)

	// While loop
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("sum 2:", sum)

	// Forever, break and continue
	n := 0
	for {
		n += 1
		if n > 10 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	var sequence = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	// range function can return index and item in a for loop
	for i, a := range sequence {
		fmt.Println(i, a)
	}
}
