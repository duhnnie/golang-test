package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x + v2.x, y: v1.y + v2.y }
}

func (v1 Vector2D) Substract(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x - v2.x, y: v1.y - v2.y }
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{ x: v1.x * v2.x, y: v1.y * v2.y }
}

func (v1 Vector2D) AddV(n float64) Vector2D {
	return Vector2D{ x: v1.x + n, y: v1.y + n }
}

func (v1 Vector2D) MultiplyV(n float64) Vector2D {
	return Vector2D{ x: v1.x * n, y: v1.y * n }
}

func (v1 Vector2D) DivideV(n float64) Vector2D {
	return Vector2D{ x: v1.x / n, y: v1.y / n }
}

func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{
		x: math.Max(math.Min(v1.x, upper), lower),
		y: math.Max(math.Min(v1.y, upper), lower),
	}
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(
		math.Pow(v1.x - v2.x, 2) + math.Pow(v1.y - v2.y, 2),
	)
}