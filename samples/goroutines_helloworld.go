// goroutines_helloworld
package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func say(s string) {
	time.Sleep(1 * time.Second)
	elapsed := time.Now().Sub(start).Round(time.Second)
	fmt.Println(elapsed, s)
}

func main() {
	say("hello")
	say("world")
	println("========")
	go say("HELLO")
	say("WORLD")
	time.Sleep(1 * time.Second)
}
