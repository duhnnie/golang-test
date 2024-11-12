package main

import "fmt"

func init() {
	fmt.Println("Hello from init function")
}

func main() {
	fmt.Println("Hello from main function")
	p := new(Person)
	fmt.Println(p)
}
