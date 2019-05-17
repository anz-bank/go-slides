package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func easyErrHandling() {

	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// NOTE: scanner.Scan() returns a boolean as opposed to an error
	// This makes the looping easier
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
	}
	// Output:
	// 1234
	// 5678
	// Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range
}

func main() {
	easyErrHandling()
}
