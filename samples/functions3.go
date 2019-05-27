package main

import (
	"fmt"
)

func filter(arr []int, pred func(int) bool) []int {
	filtered := make([]int, 0, len(arr)) // make new slice with max capacity
	for _, val := range arr {
		if pred(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered[:len(filtered)]
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// You can pass functions inline
	even := filter(vals, func(val int) bool { return val%2 == 0 })
	fmt.Println(even)

	// you can assign functions to variables
	mul3Pred := func(val int) bool {
		return val%3 == 0
	}
	mul3 := filter(vals, mul3Pred)
	fmt.Println(mul3)

	// inline function declarations are closures, they capture surrounding scope
	mul2and3Pred := func(val int) bool {
		return val%2 == 0 && mul3Pred(val) // mul3Pred captured from surrounding scope
	}
	mul6 := filter(vals, mul2and3Pred)
	fmt.Println(mul6)
}
