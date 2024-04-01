package main

import "fmt"

// declare a struct with two propeties: x: int, y: int
type Vertex struct {
	X int
	Y int
}

func main() {
	// create a new Vertex
	v := Vertex{1, 2}
	// assing a value to one of its fields
	v.Y = 4
	// print the vertex field
	fmt.Println(v.Y)
}