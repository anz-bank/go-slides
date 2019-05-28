// json - zero values
package main

import (
	"encoding/json"
	"fmt"
)

// Using primitive types
type Person struct {
	Weight int `json:"height"`
	Age    int `json:"age"`
}

// Using pointer types
type PointerPerson struct {
	Weight *int `json:"height"`
	Age    *int `json:"age"`
}

func main() {
	basic := Person{Age: 0}
	bytes, _ := json.Marshal(basic)
	fmt.Println(string(bytes))

	age := 0
	pointer := PointerPerson{Age: &age}
	bytes, _ = json.Marshal(pointer)
	fmt.Println(string(bytes))
}
