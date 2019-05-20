// Maps
package main

import "fmt"

func main() {
	// A map maps keys to values.
	// Create an empty map with builtin `make`:
	m1 := make(map[string]int)

	m1["Australia"] = 25
	m1["Austria"] = 9
	m1["India"] = 1340
	val := m1["Australia"]
	fmt.Println("map:", m1)
	fmt.Println("val:", val)
	fmt.Println("len:", len(m1))

	delete(m1, "Austria")
	fmt.Println("map:", m1)

	// Initialize a new map with a map literal
	m2 := map[string]string{"Melbourne": "VIC", "Sydney": "NSW"}
	fmt.Println("m2 :", m2)

	// An empty map literal is a substitute for make
	m3 := map[string]string{}
	fmt.Println("m3 :", m3, make(map[string]string))

	v, ok := m2["Brisbane"]
	fmt.Printf(`Value: "%s" Present?: %v`+"\n", v, ok)
}
