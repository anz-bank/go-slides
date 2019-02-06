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
	val := m1["India"]
	fmt.Println("map:", m1)
	fmt.Println("val:", val)
	fmt.Println("len:", len(m1))

	delete(m1, "Austria")
	fmt.Println("map:", m1)

	// Initialize a new map with a map literal
	m2 := map[string]string{"welcome": "ಸ್ವಾಗತ", "friend": "ಸ್ನೇಹಿತ"}
	fmt.Println("m2 :", m2)

	v, ok := m2["World peace"]
	fmt.Printf(`Value: "%s" Present?: %v`+"\n", v, ok)
}
