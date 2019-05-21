package main

import (
	"errors"
	"fmt"
)

var primes []int
var ErrBufferOverflow = errors.New("buffer overflow")

type Writer struct {
	buf []int
}

type errWriter struct {
	w   *Writer
	err error
}

func (w *Writer) writePrimes(b []int) (int, error) {
	if len(b)+len(w.buf) >= 5 {
		return len(w.buf), ErrBufferOverflow
	}
	w.buf = append(w.buf, b...)
	return len(w.buf), nil
}

func (ew *errWriter) writePrimes(b []int) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.writePrimes(b)
}

func writeCheck(ew *errWriter) {
	ew.writePrimes(primes[0:2])
	ew.writePrimes(primes[2:5])
	if ew.err != nil {
		fmt.Println("writePrimes:", ew.err)
		return
	}
}

func main() {
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	ew := &errWriter{
		w: &Writer{
			buf: make([]int, 0),
		},
		err: nil,
	}
	writeCheck(ew)
}
