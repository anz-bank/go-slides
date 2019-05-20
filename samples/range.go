// Range
package main

import "fmt"

func main() {
	// The `range` form of the for loop iterates over a slice or map.

	// In the case of slices, `range` returns two values for each iteration:
	// index and copy of element at index (i, val).
	vals := []int{1, 2, 3}
	for i, val := range vals {
		fmt.Printf("index: %d, value: %d\n", i, val)
	}

	// Ignore the slice index with the blank identifier `_`
	sum := 0
	for _, val := range vals {
		sum += val
	}
	fmt.Println("sum:", sum)

	// `range` on map iterates over key/value pairs
	m := map[string]string{"Melbourne": "VIC", "Sydney": "NSW"}
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range m {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code points
	for i, c := range "we🖤u" {
		fmt.Printf("character %c starts at byte index %d\n", c, i)
	}
}
