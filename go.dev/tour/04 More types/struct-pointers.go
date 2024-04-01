package main

import "fmt"

// declare a struct with two propeties: x: int, y: int
type Vertex struct {
	X int
	Y int
}

func main() {
	// create a new Vertex
	v := Vertex{3, 4}
	// create a pointer to the vertex
	p := &v
	// using the pointer, assing a value to one of its fields
	p.X = 5 // Also (*p).X = 5 is valid
	// print the vertex field
	fmt.Println(p.X)
}