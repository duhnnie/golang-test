package main

import "fmt"

type Person struct {
	name string
	age  int
}

func init() {
	fmt.Println("hello from init in person.go")
}
