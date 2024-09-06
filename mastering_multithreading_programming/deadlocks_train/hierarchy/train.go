package hierarchy

import (
	"example/deadlocks_train/common"
	"sort"
	"time"
)

func lockIntersectionsInDistance(train *common.Train, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection

	for _, crossing := range crossings {
		if (reserveStart <= crossing.Position && reserveEnd >= crossing.Position && crossing.Intersection.LockedBy != train.Id) {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func (intA, intB int) bool {
		return intersectionsToLock[intA].Id < intersectionsToLock[intB].Id
	})

	for _, intersection := range intersectionsToLock {
		intersection.Mutex.Lock()
		intersection.LockedBy = train.Id

		time.Sleep(10 * time.Millisecond)
	}
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		for _, crossing := range crossings {
			if train.Front + 1 == crossing.Position {
				lockIntersectionsInDistance(train, crossing.Position, crossing.Position + train.Length, crossings)
			}

			trainBack := train.Front - train.Length

			if trainBack == crossing.Position + 1 {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}

		train.Front += 1;
		time.Sleep(30 * time.Millisecond)
	}
}