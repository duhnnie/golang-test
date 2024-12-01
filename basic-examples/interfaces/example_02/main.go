package main

import "fmt"

type LivingThing interface {
	grow()
}

type Animal interface {
	makeNoise()
	grow()
}

type Dog struct{}

func (d Dog) makeNoise() {
	fmt.Println("woof!")
}

func (d Dog) grow() {
	fmt.Println("I'm a dog and I'd just grown")
}

func (d Dog) sitDown() {
	fmt.Println("I'm a dog and I'm sitted down")
}

func getLivingThing() LivingThing {
	return &Dog{}
}

func main() {
	// We have a LivingThing interface value,
	// so we can call only the methods defined for that interface.
	lt := getLivingThing()
	lt.grow()

	// Animal is another interface that satisfies LivingThing, but it
	// defines another method: makeNoise(), so we try to obtain Animal from LivingThing
	animal, ok := lt.(Animal)

	if !ok {
		panic("lt is not Animal")
	}

	animal.makeNoise()

	// We try to get the concrete Dog type from animal, that's the only way to call sitDown()
	dog, ok := animal.(*Dog)

	if !ok {
		panic("animal is not a Dog")
	}

	dog.sitDown()
}
