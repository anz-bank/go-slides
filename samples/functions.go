// Functions
package main

import "fmt"

func myWorld() string {
	return "Vienna, Melbourne, London, Jamnagar and Nairobi"
}

func main() {
	fmt.Println("Hello, Melbourne")
	fmt.Println("Hello", myWorld())
}
