// errors
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Somthing's fishy")
	fmt.Println(err)
	err = fmt.Errorf("Somthing's fishy after %d retries", 256)
	fmt.Println(err)
}
