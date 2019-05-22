// channels_range
package main

import (
	"fmt"
	"strings"
	"time"
)

func speak(ch chan string) {
	text := `We hold these truths to be self-evident, that all men are created equal, that they are endowed by their creator with certain unalienable rights, that among these are life, liberty, and the pursuit of happiness.`
	parts := strings.Split(text, ",")
	for _, part := range parts {
		ch <- strings.Trim(part, " .")
		time.Sleep(2 * time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go speak(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
