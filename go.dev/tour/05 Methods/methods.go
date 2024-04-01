package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method on a struct
func (v Vertex) Abs() float64 {
	return math.Abs(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64

// method on MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	vertex := Vertex{ 3, 4 }
	fmt.Println(vertex.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
