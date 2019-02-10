// channels
package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	fmt.Println("ping:", "ping received")
	time.Sleep(2 * time.Second)
	fmt.Println("ping:", "sending pong")
	ch <- "PONG"
}

func main() {

	// A Channel is a typed conduit through which you can send and receive values
	// with the channel operator, <-
	// Create a channel with make
	ch := make(chan string)
	go ping(ch)
	fmt.Println("main: waiting for ping response")
	str := <-ch
	fmt.Println("main: received", str)
}
