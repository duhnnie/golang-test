package deadlock

import (
	"example/deadlocks_train/common"
	"time"
)

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		for _, crossing := range crossings {
			if train.Front -1 == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.Id
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