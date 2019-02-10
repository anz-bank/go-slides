// select
package main

import (
	"fmt"
	"time"
)

func bomb1() {
	tick := time.Tick(400 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tick ")
		case <-boom:
			fmt.Print("BOOM💥")
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Print(". ")
		}
	}
}

func bomb2() {
	tick := time.Tick(400 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tock ")
		case <-boom:
			fmt.Print("BOOM2 💣")
			return
		default:
			fmt.Print(". ")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	bomb1()
	fmt.Println()
	bomb2()
}
