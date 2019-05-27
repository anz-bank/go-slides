package main

import "fmt"

// Aliasing a type allows type checking
type ErrCode uint8

// Declaring a set of constants effectively creates an enum of that type
const (
	ErrInvalidInput ErrCode = iota
	ErrDuplicate
	ErrNotFound
)

// Aliased types can be given methods too
func (e ErrCode) String() string {
	switch e {
	case ErrInvalidInput:
		return "invalid input"
	case ErrDuplicate:
		return "duplicate"
	case ErrNotFound:
		return "not found"
	}
	return "unknown"
}

// Implement the error interface
func (e ErrCode) Error() string {
	return e.String()
}

func printCode(code ErrCode) {
	// What is printed depends on print directive
	fmt.Printf("Val: %d\nString: %s\n\n", code, code)
}

func main() {
	// Enum constants can be passed in
	printCode(ErrInvalidInput)

	// So can raw values, they may not be a predefined value
	printCode(127)

	// but variables with type uint8 cannot
	var c uint8
	c = 1
	// printCode(c) // Compile error, invalid type

	// However you can typecast it to your enum type
	printCode(ErrCode(c))
}
