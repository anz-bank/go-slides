package main

import (
	"fmt"
	"strconv"
)

type MyStruct struct {
	n int
}

func (m MyStruct) String() string {
	return strconv.Itoa(m.n)
}

func RawLoop(arr []MyStruct) {
	// A loop over structs COPIES each array element into item
	for _, item := range arr {
		item.n = 1
	}
}

func PointerLoop(arr []*MyStruct) {
	// Looping over pointers copies the POINTER, which still points to
	// the original object
	for _, item := range arr {
		item.n = 1
	}
}

func main() {
	// Looping over raw struct will not allow you to modify them
	arr := []MyStruct{{0}, {0}, {0}}
	RawLoop(arr)
	fmt.Println(arr)

	// Looping over pointers fixes this
	pointerArr := []*MyStruct{{0}, {0}, {0}}
	PointerLoop(pointerArr)
	fmt.Println(pointerArr)
}
