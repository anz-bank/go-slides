package main

import (
	"errors"
	"fmt"
)

type Writer struct {
	buf []int
}

var primes []int
var ErrBufferOverflow = errors.New("buffer overflow")

func (w *Writer) writePrimes(b []int) (int, error) {
	if len(b)+len(w.buf) >= 5 {
		return len(w.buf), ErrBufferOverflow
	}
	w.buf = append(w.buf, b...)
	return len(w.buf), nil
}

func writeCheck(w *Writer) {
	_, err := w.writePrimes(primes[0:2])
	if err != nil {
		fmt.Println("writePrimes:", err)
		return
	}
	_, err = w.writePrimes(primes[2:5])
	if err != nil {
		fmt.Println("writePrimes:", err)
		return
	}
}

func main() {
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	w := &Writer{
		buf: make([]int, 0),
	}
	writeCheck(w)
}
