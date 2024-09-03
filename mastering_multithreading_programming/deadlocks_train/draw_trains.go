package main

import (
	"example/deadlocks_train/common"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var colors = [4]color.RGBA {
	color.RGBA{233, 33, 40, 255},
	color.RGBA{78, 151, 210, 255},
	color.RGBA{251, 170, 26, 255},
	color.RGBA{11, 132, 54, 255},
}

var white = color.RGBA{185, 185, 185, 255}

func DrawIntersections(screen *ebiten.Image) {
	drawIntersection(screen, intersections[0], 145, 145)
	drawIntersection(screen, intersections[1], 175, 145)
	drawIntersection(screen, intersections[2], 175, 175)
	drawIntersection(screen, intersections[3], 145, 175)
}

func DrawTrains(screen *ebiten.Image) {
	drawXTrain(screen, trains[0], 1, 10, 135)
	drawYTrain(screen, trains[1], 1, 10, 185)
	drawXTrain(screen, trains[2], -1, 310, 185)
	drawYTrain(screen, trains[3], -1, 310, 135)
}

func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < 300; i++ {
		screen.Set(10 + i, 135, white)
		screen.Set(185, 10 + i, white)
		screen.Set(310 - i, 185, white)
		screen.Set(135, 310  - i, white)
	}
}

func drawIntersection(screen *ebiten.Image, intersection *common.Intersection, x int, y int) {
	color := white

	if intersection.LockedBy >= 0 {
		color = colors[intersection.LockedBy]
	}

	screen.Set(x, y, color)
	screen.Set(x - 1, y, color)
	screen.Set(x + 1, y, color)
	screen.Set(x, y - 1, color)
	screen.Set(x, y + 1, color)
}

func drawXTrain(screen *ebiten.Image, train *common.Train, dir int, start int, yPos int) {
	s := start + (dir * (train.Front - train.Length))
	e := start + (dir * train.Front)

	for i := math.Min(float64(s), float64(e)); i < math.Max(float64(s), float64(e)); i++ {
		screen.Set(int(i) - dir, yPos - 1, colors[train.Id])
		screen.Set(int(i), yPos, colors[train.Id])
		screen.Set(int(i) - dir, yPos + 1, colors[train.Id])
	}
}

func drawYTrain(screen *ebiten.Image, train *common.Train, dir int, start int, xPos int) {
	s := start + (dir * (train.Front - train.Length))
	e := start + (dir * train.Front)

	for i := math.Min(float64(s), float64(e)); i < math.Max(float64(s), float64(e)); i ++ {
		screen.Set(xPos, int(i), colors[train.Id])
		screen.Set(xPos - 1, int(i) - dir, colors[train.Id])
		screen.Set(xPos + 1, int(i) - dir, colors[train.Id])
	}
}
