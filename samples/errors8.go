package main

import (
	"errors"
	"fmt"
	"net"
)

type BadRequestError string

func (e BadRequestError) Error() string {
	return string(e)
}

var ErrAbortHandler = errors.New("net/http: abort handler")

func printError(err error) {

	switch err.(type) {
	case BadRequestError:
		fmt.Println("BadRequestError: ", err)
	case *net.DNSError:
		fmt.Println("DNSError: ", err)
	case error:
		fmt.Println("Generic: ", err)
	default:
		fmt.Println("Unknown error: ", err)
	}
}

func main() {
	var e BadRequestError = "400 bad request"
	dnsErr := net.DNSError{
		Err:         "Unknown host",
		Name:        "bad.host.xyz",
		IsTimeout:   false,
		IsTemporary: false,
	}

	printError(e)
	printError(&dnsErr)
	printError(ErrAbortHandler)
}
