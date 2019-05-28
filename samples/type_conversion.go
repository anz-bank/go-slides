// Type conversion
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2
	f := math.Sqrt(float64(x))
	z := uint(f)
	fmt.Printf("x: %v %T\n", x, x)
	fmt.Printf("f: %v %T\n", f, f)
	fmt.Printf("z: %v %T\n", z, z)
}
