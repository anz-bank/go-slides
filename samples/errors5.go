package main

import (
	"bufio"
	"fmt"
	"os"
)

// DOES NOT COMPILE.
// Only shows usage
func main() {

	// call to scan the input using Scan method.
	// Note that it returns a bool.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println(token)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error processing tokens: ", err)
	}

	// While it could have been more tedious if Scan
	// returned an error instead of a bool
	// The caller would then have to write
	scanner := bufio.NewScanner(input)
	for {
		token, err := scanner.Scan()
		if err != nil {
			return err // or maybe break
		}
		// process token
	}
}
