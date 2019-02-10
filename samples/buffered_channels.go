// Buffered channels
package main

import "fmt"

func main() {
	// With buffered channels, provide the buffer length as the second argument to `make`
	// `Send` only blocks when the buffer is full, `Receive` only blocks when the buffer is empty
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
