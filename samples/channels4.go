package main

import (
	"fmt"
	"time"
)

// Why is the output of channelWrite in order
// Try this exercise on your laptop.
// Set GOMAXPROCS environment variable to 1
// and a higher value to experiment.

var nWrites int = 6

func channelWrite(ch chan int, value int) {
	ch <- value
	fmt.Printf("channelWrite[%d]: [%T] <- %v\n", value, ch, value)
}

func main() {
	ch := make(chan int, 5)
	for i := 0; i < nWrites; i++ {
		go channelWrite(ch, i)
	}
	time.Sleep(5 * time.Second)
	// Uncomment the line below to see the behavior of the code
	// on your machine and playground.
	// close(ch)
	fmt.Println("main: exiting")
}
