// Interfaces
package main

import "fmt"

type scaler interface {
	scale(float64)
}

func main() {
	var s scaler
	s = &rect{2, 3}
	s.scale(2.0)
	fmt.Println(s)

	s = &circle{20}
	s.scale(.5)
	fmt.Println(s)

	// Type assertions
	c1 := s.(*circle)
	c2, ok := s.(*circle)
	fmt.Println("Circle?", ok, c1, c2)
	r1, ok := s.(*rect)
	fmt.Println("Rect?", ok, r1)
	// r2 := s.(rect) // panics
}

type rect struct {
	a, b int
}

func (r *rect) scale(s float64) {
	r.a = int(float64(r.a) * s)
	r.b = int(float64(r.b) * s)
}

type circle struct {
	r int
}

func (c *circle) scale(s float64) {
	c.r = int(float64(c.r) * s)
}
