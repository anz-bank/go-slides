package main

import "fmt"

func fibGenerator(n int, ch chan int) {
	i, j := 0, 1
	for k := 0; k < n; k++ {
		ch <- i
		i, j = j, i+j
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibGenerator(cap(ch), ch)
	fmt.Println("Fibonacci Series:")
	for k := range ch {
		fmt.Println(k)
	}
}
