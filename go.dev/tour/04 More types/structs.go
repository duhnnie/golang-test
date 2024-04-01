package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertex := Vertex{1, 2}

	fmt.Printf("%+v\n", vertex)
}