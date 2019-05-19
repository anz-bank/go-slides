package main

import (
	"fmt"
	"time"
)

var nWrites int = 1

func channelWrite(ch chan int, value int) {
	ch <- value
	fmt.Printf("channelWrite[%d]: [%T] <- %v\n", value, ch, value)
}

func main() {
	ch := make(chan int)
	for i := 0; i < nWrites; i++ {
		go channelWrite(ch, i)
	}
	time.Sleep(5 * time.Second)
	// Uncomment the line below to see the behavior of the code
	// on your machine and playground.
	// close(ch)
	fmt.Println("main: exiting")
}
