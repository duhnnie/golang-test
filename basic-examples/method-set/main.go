package main

import "fmt"

type Talkable interface {
	salute()
}

type Dog struct {
	name  string
	breed string
	age   int
}

func (d Dog) bark() {
	fmt.Println("woof")
}

func (d *Dog) salute() {
	fmt.Printf("Woff! My name is %s\n", d.name)
}

func introduce(i Talkable) {
	i.salute()
}

func main() {
	dog := Dog{
		name:  "Gandhi",
		breed: "Husky",
		age:   2,
	}

	dog.bark()

	introduce(&dog)

	Dog.bark(dog)
	(*Dog).salute(&dog)
}
