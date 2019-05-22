package main

import (
	"fmt"
)

type Foo struct {
	n int
}

// Input variables are copied in
func RawStructModifier(data Foo) {
	// This modifies a copy, not the original
	data.n = 1
}

// Copying a pointer still gives you access to the original data
func PointerModifier(data *Foo) {
	data.n = 1
}

func main() {
	f := Foo{0}
	RawStructModifier(f)
	fmt.Println(f)
	PointerModifier(&f)
	fmt.Println(f)
}
