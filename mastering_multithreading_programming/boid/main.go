package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount = 500
	viewRadius = 13
	adjustmentRatio = 0.015
)

var (
	green = color.RGBA{ 10, 255, 50, 255 }
	boids [boidCount]*Boid
	boidsMap [screenWidth + 1][screenHeight + 1]int
)


type Game struct {}

func (g Game) Update(image *ebiten.Image) error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		if boid != nil {
			screen.Set(int(boid.position.x - 1), int(boid.position.y), green)
			screen.Set(int(boid.position.x + 1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y - 1), green)
			screen.Set(int(boid.position.x), int(boid.position.y + 1), green)
		}
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < len(boidsMap); i++ {
		for j := 0; j < len(boidsMap[i]); j++ {
			boidsMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth * 2, screenHeight * 2)
	ebiten.SetWindowTitle("Boids in a box")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}