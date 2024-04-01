package main

import "fmt"

func main() {
	// Create two integers
	i, j := 10, 173

	// create a pointer to one of the integers
	p := &i
	// print the value of the integer through the pointer
	fmt.Println(*p)
	// set a new value to the interger var through the pointer
	*p = 45
	// print the new value of the modified interger
	fmt.Println(i)

	// get the pointer of the second integer
	p = &j
	// using the pointer, modifiy the underlying value, an store the value
	*p = *p / 3
	// print the new value
	fmt.Println(j)
}