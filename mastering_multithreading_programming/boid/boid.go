package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id int
}

func (b *Boid) calculateAcceleration() Vector2D {
	lower, upper := b.position.AddV(-viewRadius), b.position.AddV(viewRadius)
	avgVelocity := Vector2D{}
	count := 0.0

	for i := int(math.Max(lower.x, 0)); i < int(math.Min(upper.x, screenWidth)); i ++ {
		for j := int(math.Max(lower.y, 0)); j < int(math.Min(upper.y, screenHeight)); j ++ {
			if otherBoidId := boidsMap[i][j]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					count++
				}
			}
		}
	}

	acceleration := Vector2D{x: 0, y: 0}

	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		acceleration = avgVelocity.Substract(b.velocity).MultiplyV(adjustmentRatio)
	}

	return acceleration
}

func (b *Boid) moveOne() {
	acc := b.calculateAcceleration()
	newVelocity := b.velocity.Add(acc).limit(-1, 1)
	b.velocity = newVelocity

	boidsMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidsMap[int(b.position.x)][int(b.position.y)] = b.id

	next := b.position.Add(b.velocity)

	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D { x: -b.velocity.x, y: b.velocity.y }
	}

	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{ x: b.velocity.x, y: -b.velocity.y }
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	boid := Boid{
		position: Vector2D{
			x: rand.Float64() * screenWidth,
			y: rand.Float64() * screenHeight,
		},
		velocity: Vector2D{
			x: (rand.Float64() * 2) - 1,
			y: (rand.Float64() * 2) - 1,
		},
		id: bid,
	}

	boids[bid] = &boid
	boidsMap[int(boid.position.x)][int(boid.position.y)] = bid
	go boid.start()
}