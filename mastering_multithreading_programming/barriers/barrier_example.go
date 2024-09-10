package main

import (
	"fmt"
	"time"
)

func waitOnBarrier(name string, timeToSleep time.Duration, barrier *Barrier) {
	for {
		fmt.Printf("%v is running...\n", name)
		time.Sleep(timeToSleep * time.Second)
		fmt.Printf("%v is waiting...\n", name)
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)
	time.Sleep(100 * time.Second)
}