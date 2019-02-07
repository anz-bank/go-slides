// Functions
package main

import "fmt"

// Multiple parameters of same type shorthand
// equivalent to `func sum(a int, b int) int`
func sum(a, b int) int {
	return a + b
}

// Multiple return values
func gimmeTwo() (string, int) {
	return "0b10", 2
}

// Named return values, use with care
func location(city string) (lat, lng float64) {
	switch city {
	case "Bangalore", "Bangaluru":
		lat, lng = 12.9716, 77.5946
	case "Melbourne":
		lat, lng = -37.8136, 144.9631
	}
	return
}

func main() {
	fmt.Println("sum(2,5)", sum(2, 5))
	fmt.Println(location("Bangalore"))
	fmt.Println(gimmeTwo())
}
