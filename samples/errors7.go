package main

import (
	"errors"
	"fmt"
)

func writeCheck() {

	var err error

	buf := make([]byte, 0, 0)
	write := func(n byte) {
		if len(buf) == 50 {
			err = errors.New("buffer overflow")
			return
		}
		buf = append(buf, n)
	}
	for i := 0; i < 100; i++ {
		write(byte(i))
	}
	if err != nil {
		fmt.Println("write:", err)
		return
	}
}

func main() {
	writeCheck()
}
