// Values
package main

import "fmt"

func main() {

	fmt.Printf("High five\n")        // interpreted string literal
	fmt.Println(`He says, "Run!"\n`) // raw string literal

	fmt.Println("13 + 8 =", 13+8)               // ints
	fmt.Println("333.0 / 106.0 =", 333.0/106.0) // floats

	fmt.Println("and", true && false) // bool
	fmt.Println("not", !true)
}
