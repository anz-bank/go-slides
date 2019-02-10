// goroutines
package main

import (
	"fmt"
	"time"
)

func sleepyPrint(i interface{}) {
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}

func main() {
	for i := 0; i < 3; i++ {
		sleepyPrint(i)
	}
	sleepyPrint("Almost finished")
	for i := 0; i < 3; i++ {
		go sleepyPrint(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}
