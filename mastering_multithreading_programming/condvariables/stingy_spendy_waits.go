package main

import (
	"fmt"
	"sync"
	"time"
)

var money = 100
var lock = sync.Mutex{}
var moneyDeposited = sync.NewCond(&lock)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money = money + 10
		fmt.Println("Stingy sees balance:", money)
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Stingy done!")
}

func spendy() {
	for i:= 1; i <= 1000; i++ {
		lock.Lock()
		for money - 20 < 0 {
			moneyDeposited.Wait()
		}
		money = money - 20
		fmt.Println("Spendy sees balance:", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Spendy done!")
}

func main() {
	fmt.Println("Starting...")
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}