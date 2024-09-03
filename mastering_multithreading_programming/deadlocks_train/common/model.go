package common

import "sync"

type Train struct {
	Id int
	Length int
	Front int
}

type Intersection struct {
	Id int
	Mutex sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position int
	Intersection *Intersection
}

func (c *Crossing) IsOverCrossing(train *Train, extra int) bool {
	back := (train.Front + extra) - train.Length

	return back <= c.Position && train.Front + extra >= c.Position
}
