package main

import "fmt"

type Writer interface {
	Write()
}

type Singer interface {
	Sing()
}

type WriterSinger interface {
	Writer
	Singer
}

type Artist struct {
	Name string
}

func (a *Artist) Write() {
	fmt.Printf("📝My name is %s\n", a.Name)
}

func (a *Artist) Sing() {
	fmt.Printf("🎵%s, my name\n", a.Name)
}

func main() {
	var bob WriterSinger
	bob = &Artist{
		Name: "Bob Dylan",
	}

	bob.Write()
	bob.Sing()
}
