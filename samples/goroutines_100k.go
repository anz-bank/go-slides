package main

import (
	"fmt"
	"runtime"
	"time"
)

var nGoroutines int

func memFunc() {
	b := make([]int, 1024)
	for i := 0; i < len(b); i++ {
		b[0] = i
		time.Sleep(1 * time.Microsecond)
	}
}

func report() {
	for {
		n := runtime.NumGoroutine()
		if n > 1 && nGoroutines != n {
			nGoroutines = n
			fmt.Printf("\r                                    ")
			fmt.Printf("\r# Goroutines: %d/100,000", n)
			time.Sleep(1 * time.Millisecond)
		} else if n == 1 {
			fmt.Printf("\n# Goroutines: %d/100,000 (Exiting)\n", n)
			break
		}
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		if i%10000 == 0 {
			fmt.Printf("Started %d goroutines\n", i)
		}
		go memFunc()
	}
	report()
}
