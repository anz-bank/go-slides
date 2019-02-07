// Methods
package main

import "fmt"

type rect struct {
	a, b int
}

// A method is a function with a special receiver argument.
// The `area` method has a value receiver of type `rect`
func (r rect) area() int {
	return r.a * r.b
}

// Methods with pointer receivers (`*rect`) can modify the value to which the receiver points
func (r *rect) scale(s float64) {
	r.a = int(float64(r.a) * s)
	r.b = int(float64(r.b) * s)
}

func main() {
	r := rect{a: 4, b: 6}

	fmt.Println("rect: ", r)
	fmt.Println("area: ", r.area())
	r.scale(.5)
	fmt.Println("rect: ", r)

	rp := &r
	// As a convenience, Go interprets rp.area() as (&rp).area()
	fmt.Println("area: ", rp.area())
}
