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

var abortErr = errors.New("net/http: abort")

func printError(err error) {
	switch err.(type) {
	case BadRequestError:
		fmt.Println("BadRequestError: ", err)
	case *net.DNSError:
		fmt.Println("DNSError: ", err)
	default:
		fmt.Println("error: ", err)
	}
}

func main() {
	var badRequestErr BadRequestError = "400 bad request"
	dnsErr := &net.DNSError{Err: "Unknown host"}

	printError(badRequestErr)
	printError(dnsErr)
	printError(abortErr)
}
