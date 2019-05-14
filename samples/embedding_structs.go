package main

import "fmt"

type Pet struct {
	Name string
}

type Cat struct {
	Pet
	Lives uint32
}

func main() {
	cat := Cat{
		Pet:   Pet{Name: "Charlotte"},
		Lives: 9,
	}

	fmt.Printf("My name is %s \n", cat.Name)
	fmt.Printf("I have %d lives \n", cat.Lives)
}
