package main

// Pointers
import "fmt"

type MyStruct struct {
	n int
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
	for _, item := range arr {
		fmt.Println(item.n)
	}

	fmt.Println()

	// Looping over pointers fixes this
	pointerArr := []*MyStruct{{0}, {0}, {0}}
	PointerLoop(pointerArr)
	for _, item := range pointerArr {
		fmt.Println(item.n)
	}
}
