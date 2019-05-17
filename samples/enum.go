// Enum
package main

import "fmt"

// Aliasing a type allows type checking
type ErrorEnum string

// Implement the error interface
func (e ErrorEnum) Error() string {
	return string(e)
}

const (
	ErrInvalidArgument ErrorEnum = "Invalid value"
	ErrNotFound        ErrorEnum = "Not found"
	ErrUnknown         ErrorEnum = "Unknown"
)

func isEnum(err ErrorEnum) {
	fmt.Printf("%s is and ErrorEnum\n", err)
}

func main() {
	// Enum constants can be passed in
	isEnum(ErrInvalidArgument)

	// So can raw strings
	isEnum("WAT")

	// but variables with type string cannot
	var s string
	s = "Hello"
	// isEnum(s) // Compile error, invalid type

	// However you can typecast it to your enum type
	isEnum(ErrorEnum(s))
}
