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

func (b *Boid) borderBounce(pos, maxPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxPos - viewRadius {
		return 1 / pos - maxPos
	}

	return 0
}

func (b *Boid) calculateAcceleration() Vector2D {
	lower, upper := b.position.AddV(-viewRadius), b.position.AddV(viewRadius)
	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0

	rwLock.RLock()
	for i := int(math.Max(lower.x, 0)); i < int(math.Min(upper.x, screenWidth)); i ++ {
		for j := int(math.Max(lower.y, 0)); j < int(math.Min(upper.y, screenHeight)); j ++ {
			if otherBoidId := boidsMap[i][j]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Substract(boids[otherBoidId].position).DivideV(dist))
					count++
				}
			}
		}
	}
	rwLock.RUnlock()

	acceleration := Vector2D{
		x: b.borderBounce(b.position.x, screenWidth),
		y: b.borderBounce(b.position.y, screenHeight),
	}

	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		avgPosition = avgPosition.DivideV(count)
		accelerationAlignment := avgVelocity.Substract(b.velocity).MultiplyV(adjusmentRate)
		accelerationCohesion := avgPosition.Substract(b.position).MultiplyV(adjusmentRate)
		accelerationSeparation := separation.MultiplyV(adjusmentRate)
		acceleration = acceleration.Add(accelerationAlignment).Add(accelerationCohesion).Add(accelerationSeparation)
	}

	return acceleration
}

func (b *Boid) moveOne() {
	acc := b.calculateAcceleration()

	rwLock.Lock()
	newVelocity := b.velocity.Add(acc).limit(-1, 1)
	b.velocity = newVelocity

	boidsMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidsMap[int(b.position.x)][int(b.position.y)] = b.id
	rwLock.Unlock()
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