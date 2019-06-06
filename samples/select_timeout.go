// select_timeout
package main

import (
	"fmt"
	"os"
	"time"
)

func write(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int)
	go write(ch)
	t := time.Tick(3 * time.Second)
	for {
		select {
		case x := <-ch:
			fmt.Println(time.Now(), x)
		case <-t:
			fmt.Println(time.Now(), "Timed out")
			os.Exit(1)
		}
	}
}
