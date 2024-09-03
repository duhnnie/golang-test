package main

import (
	"example/deadlocks_train/common"
	"example/deadlocks_train/deadlock"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return 320, 320
}

const trainLength = 100

var intersections [4]*common.Intersection
var trains [4]*common.Train
var crossings [8]*common.Crossing

func main() {
	ebiten.SetWindowSize(320 * 3, 320 * 3)
	ebiten.SetWindowTitle("Trains in a box")

	// rand.Seed(uint64(time.Now().UnixNano()))

	for i := 0; i < 4; i++ {
		trains[i] = &common.Train{
			Id: i,
			Length: trainLength, // rand.Intn(101) + 30,
			Front: 0,
		}

		intersections[i] = &common.Intersection{
			Id: i,
			Mutex: sync.Mutex{},
			LockedBy: -1,
		}
	}

	for i := 0; i < 4; i++ {
		crossings[i * 2] = &common.Crossing{
			Position: 125,
			Intersection: intersections[i],
		}

		crossings[(i * 2) + 1] = &common.Crossing{
			Position: 175,
			Intersection: intersections[(i + 1) % 4],
		}
	}

	go deadlock.MoveTrain(trains[0], 300, []*common.Crossing{ crossings[0], crossings[1] })
	go deadlock.MoveTrain(trains[1], 300, []*common.Crossing{ crossings[2], crossings[3] })
	go deadlock.MoveTrain(trains[2], 300, []*common.Crossing{ crossings[4], crossings[5] })
	go deadlock.MoveTrain(trains[3], 300, []*common.Crossing{ crossings[6], crossings[7] })

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}