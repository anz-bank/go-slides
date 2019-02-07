// Structs
package main

import "fmt"

// This `person` struct type has `name` and `good` fields.
type person struct {
	name      string
	malicious bool
}

func main() {
	fmt.Println(person{"Bob", false})
	fmt.Println(person{name: "Alice", malicious: false})
	fmt.Println(person{name: "Trent"})

	// Use `&` for a pointer to the struct
	fmt.Println(&person{name: "Eve", malicious: true})

	// Access struct fields with a dot
	m := person{name: "Mallory", malicious: true}
	fmt.Println(m.name, m.malicious)

	// Pointers are automatically dereferenced
	mp := &m
	fmt.Println(mp.name)

	// Structs are mutable
	mp.name = "Mallet"
	fmt.Println(mp.name, m.name)
}
