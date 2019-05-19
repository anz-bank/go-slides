package main

import "fmt"

func fibGenerator(ch chan int) {
	i, j := 0, 1
	for {
		ch <- i
		i, j = j, i+j
	}
}

func main() {
	ch := make(chan int)
	go fibGenerator(ch)
	fmt.Println("Fibonacci Series:")
	for k := 0; k < 10; k++ {
		fmt.Println(<-ch)
	}
}
