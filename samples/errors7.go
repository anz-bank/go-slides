package main

import (
	"errors"
	"fmt"
)

type BadRequestError string

var ErrAbortHandler = errors.New("net/http: abort handler")

type Error interface {
	error
	Timeout() bool
	Temporary() bool
}

type DNSError struct {
	Err         string
	Name        string
	IsTimeout   bool
	IsTemporary bool
}

func (e *DNSError) Timeout() bool   { return e.IsTimeout || e.IsTemporary }
func (e *DNSError) Temporary() bool { return e.IsTemporary }
func (e DNSError) Error() string {
	return fmt.Sprintf("%s\n", e.Err)
}

func (e BadRequestError) Error() string {
	return string(e)
}

func isNetError(err error) bool {
	var ret bool

	switch err.(type) {
	case nil:
		fmt.Println("nil: no error")
		ret = false
	case BadRequestError:
		fmt.Println("BadRequestError: ", err)
		ret = true
	case DNSError:
		fmt.Println("DNSError: ", err)
	case error:
		fmt.Println("Generic: ", err)
		ret = false
	default:
		fmt.Println("Unknown error: ", err)
	}
	return ret
}

func main() {

	var e BadRequestError
	e = "Invalid JWT token"
	_ = isNetError(e)

	dnsErr := DNSError{
		Err:         "Unknown host",
		Name:        "bad.host.xyz",
		IsTimeout:   false,
		IsTemporary: false,
	}
	_ = isNetError(dnsErr)
	_ = isNetError(ErrAbortHandler)

}

