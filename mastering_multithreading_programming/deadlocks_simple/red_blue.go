package main

import (
	"fmt"
	"sync"
	"time"
)

var lock1 = sync.Mutex{}
var lock2 = sync.Mutex{}

func redRobot() {
	for {
		fmt.Println("Red: Adquiring Lock 1")
		lock1.Lock()
		fmt.Println("Red: Lock 1 adquired")
		fmt.Println("Red: Adquiring Lock 2")
		lock2.Lock()
		fmt.Println("Red: Lock 2 adquired")
		lock2.Unlock()
		lock1.Unlock()
		fmt.Println("Red: both locks released")
	}
}

func blueRobot() {
	for {
		fmt.Println("Blue: Adquiring Lock 2")
		lock2.Lock()
		fmt.Println("Blue: Lock 2 adquired")
		fmt.Println("Blue: Adquiring Lock 1")
		lock1.Lock()
		fmt.Println("Blue: Lock 1 adquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: both locks released")
	}
}

func main() {
	go redRobot()
	go blueRobot()

	time.Sleep(20 * time.Second)
}