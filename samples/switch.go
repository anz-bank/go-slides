// Switch
package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4: // comma separate expressions
		fmt.Println("three or four")
	default: // can be left out
		fmt.Println("greater than four")
	}

	fmt.Println(typeStr(true))
	fmt.Println(typeStr("bla"))
	fmt.Println(typeStr(13))
}

func typeStr(i interface{}) string {
	// Type switch
	switch t := i.(type) {
	case bool:
		return "type: bool"
	case string:
		return "type: string"
	default:
		return fmt.Sprintf("unknown type (%T)", t)
	}
}
