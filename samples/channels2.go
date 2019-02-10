// channels continued
package main

import (
	"fmt"
	"time"
)

var begin time.Time = time.Now()

func ping(ch chan string) {
	print("ping received")
	time.Sleep(2 * time.Second)
	print("sending pong")
	ch <- "PONG"
}

func main() {
	ch := make(chan string)
	go ping(ch)
	time.Sleep(1 * time.Second)
	print("WAITING")
	print(<-ch)
	resetTime()

	// Send and receive block until the other side is ready.
	// Deadlock without goroutine:
	// ping(ch)
	// print("main", <-ch)

	// Buffered channel
	ch2 := make(chan string, 1)
	ping(ch2)
	time.Sleep(1 * time.Second)
	print("WAITING") // notice the different timing
	print(<-ch2)
}

func resetTime() {
	print("-------------")
	begin = time.Now()
}

func print(i ...interface{}) {
	fmt.Print("T", int(time.Since(begin).Seconds()), " ")
	fmt.Println(i...)
}
