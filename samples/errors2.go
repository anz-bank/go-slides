// errors
package main

import (
	"fmt"
	"strconv"
)

func isInt(str string) {
	if _, err := strconv.Atoi(str); err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("looks like an int")
}

func main() {
	isInt("123")
	isInt("0x123")
}
