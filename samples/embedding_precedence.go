package main

import "fmt"

type Pet struct {
	Name string
}

type Cat struct {
	Pet
	Name string
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Pet
}

func main() {
	pet := Pet{
		Name: "Charlotte",
	}
	cat := Cat{
		Pet:  pet,
		Name: "Charly",
	}
	fmt.Printf("pet.Name:  %s\n", pet.Name)
	fmt.Printf("cat.Name:  %s\n\n", cat.Name)

	dog := Dog{
		Animal: Animal{Name: "Watchdog"},
		Pet:    Pet{Name: "Rover"},
	}
	fmt.Printf("dog.Animal.Name:  %s\n", dog.Animal.Name)
	fmt.Printf("dog.Pet.Name:     %s\n", dog.Pet.Name)
	//fmt.Printf("dog.Name:         %s\n", dog.Name) // compilation error
}
