package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Abs(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

func main() {
	vertex := Vertex{3, 4}
	fmt.Println(vertex.Abs())

	vertex.Scale(4)
	fmt.Println(vertex.Abs())
}