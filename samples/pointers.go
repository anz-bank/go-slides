// Pointers
package main

import "fmt"

func main() {
	v1 := 1
	v2 := v1
	p := &v1
	fmt.Printf("v1: %d \nv2: %d \n*p: %d \n\n", v1, v2, *p)
	*p = 2
	fmt.Printf("v1: %d \nv2: %d \n*p: %d \n\n", v1, v2, *p)
}
