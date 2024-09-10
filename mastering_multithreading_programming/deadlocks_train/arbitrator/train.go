package arbitrator

import (
	"example/deadlocks_train/common"
	"fmt"
	"sync"
	"time"
)

var locker = sync.Mutex{}
var condition = sync.NewCond(&locker)

func allFree(intersections []*common.Intersection) bool {
	for _, intersection := range intersections {
		if intersection.LockedBy >= 0 {
			return false
		}
	}

	return true
}

func lockIntersectionsInDistance(train *common.Train, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection

	for _, crossing := range crossings {
		if (reserveStart <= crossing.Position && reserveEnd >= crossing.Position && crossing.Intersection.LockedBy != train.Id) {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	locker.Lock()
	for !allFree(intersectionsToLock) {
		condition.Wait()
	}

	for _, intersection := range intersectionsToLock {
		intersection.LockedBy = train.Id
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("Train %v got 1 lock\n", train.Id)
	}
	locker.Unlock()
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		for _, crossing := range crossings {
			if train.Front + 1 == crossing.Position {
				lockIntersectionsInDistance(train, crossing.Position, crossing.Position + train.Length, crossings)
			}

			trainBack := train.Front - train.Length

			if trainBack == crossing.Position + 1 {
				locker.Lock()
				crossing.Intersection.LockedBy = -1
				condition.Broadcast()
				locker.Unlock()
			}
		}

		train.Front += 1;
		time.Sleep(30 * time.Millisecond)
	}
}