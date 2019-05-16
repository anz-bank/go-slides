package main

import (
	"bufio"
	"fmt"
	"os"
)

// If scanner.Scan() returned an error object instead
// of bool the loop would have been something like:
// for {
//     if token, err := scanner.Scan(); err != nil {
//     	   break
//     }
// }
// Every iteration would have to check for the presence
// of an error explicitly.
//
func easyErrHandling() {
	fmt.Println("Enter the input to tokenize (Type Ctrl-D to terminate):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println("> ", token)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error processing tokens: ", err)
	}
}

func main() {
	easyErrHandling()
}
