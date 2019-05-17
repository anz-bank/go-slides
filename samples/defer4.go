package main

import (
	"fmt"
	"sync"
)

func deferTest() {
	var mutex = &sync.Mutex{}

	fmt.Println("Locking resource")
	mutex.Lock()
	defer func() {
		fmt.Println("Unlocking resource")
		mutex.Unlock()
	}()

	for i := 0; i < 20; i++ {
		fmt.Printf("%v,", i)
		if i == 5 {
			fmt.Println()
			return
		}
	}
}

func main() {
	deferTest()
}
