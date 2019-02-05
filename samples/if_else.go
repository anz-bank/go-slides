// If/Else
package main

import "fmt"

func main() {

	if 117%13 == 0 {
		fmt.Println("117 is divisible by 13")
	}

	// The if statement can start with a short statement to execute before the condition
	// Variables declared inside an if short statement are also available inside any of the else blocks
	if num := 13; num < 0 {
		fmt.Println(num, "is negative")
	} else if num%2 == 0 {
		fmt.Println(num, "is positive and even")
	} else {
		fmt.Println(num, "is positive and odd")
	}
}
