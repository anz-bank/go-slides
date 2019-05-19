package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(i) // ERROR! Should be wg.Add(1) instead. Causes deadlock
		go say(fmt.Sprintf("%d: hello, world", i))
	}
	wg.Wait()
}
