package main

import (
	"fmt"
	"time"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Number of logical processors:", runtime.NumCPU())
	go say("world")
	say("hello")
}
