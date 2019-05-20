package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func splitAndParseInt(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanWords(data, atEOF)
	if err == nil && token != nil {
		_, err = strconv.ParseInt(string(token), 10, 32)
	}
	return advance, token, err
}

func easyErrHandling() {
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitAndParseInt)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s\n", err)
	}
}

func main() {
	easyErrHandling()
}
