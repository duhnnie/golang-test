package main

import "sync"

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	conditionVar sync.Cond
}

func NewBarrier(size int) *Barrier {
	mutex := &sync.Mutex{}

	return &Barrier{
		total: size,
		count: size,
		mutex: mutex,
		conditionVar: *sync.NewCond(mutex),
	}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count -= 1

	if b.count == 0 {
		b.count = b.total
		b.conditionVar.Broadcast()
	} else {
		b.conditionVar.Wait()
	}

	b.mutex.Unlock()
}