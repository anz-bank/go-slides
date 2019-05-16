// json - marshal / unmarshal
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Use `json:` tags on struct fields to customize JSON property names.
type Person struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func main() {
	bytes, _ := json.Marshal(true)
	fmt.Println(string(bytes))

	bytes, _ = json.Marshal(1)
	fmt.Println(string(bytes))

	slice := []string{"almost", "there"}
	bytes, _ = json.Marshal(slice)
	fmt.Println(string(bytes))

	person := Person{LastName: "Sting", Age: 67}
	bytes, _ = json.Marshal(person)
	fmt.Println(string(bytes))

	bytes = []byte(`{"firstName": "Julius", "lastName": "Caesar", "age": 55}`)
	if err := json.Unmarshal(bytes, &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)

	bytes = []byte(`{"name": "Mr. Shell", "game": "Shellcheckers", "fame": 100}`)
	var obj map[string]interface{}
	if err := json.Unmarshal(bytes, &obj); err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj)

}
