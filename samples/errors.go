// errors
package main

import (
	"fmt"
	"strings"
)

const (
	Invalid  = "INVALID"
	NotFound = "NOT-FOUND" // ...
)

type CodeError struct {
	Message string
	Code    string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("%s [code: %s]", e.Message, e.Code)
}

func checkTitle(str string) error {
	if strings.Title(str) != str {
		message := fmt.Sprintf(`no title case: "%s"`, str)
		return &CodeError{message, Invalid}
	}
	return nil
}

func main() {
	if err := checkTitle("John dough"); err != nil {
		fmt.Println(err)
	}
}
