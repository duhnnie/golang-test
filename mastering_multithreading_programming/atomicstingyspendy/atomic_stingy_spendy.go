package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var money int32 = 100
var waitgroup = sync.WaitGroup{}

func stingy() {
	for i := 1; i <= 100000; i++ {
		atomic.AddInt32(&money, 10)
	}

	println("Stingy done!")
	waitgroup.Done()
}

func spendy() {
	for i:= 1; i <= 100000; i++ {
		atomic.AddInt32(&money, -10)
	}

	println("Spendy done!")
	waitgroup.Done()
}

func main() {
	fmt.Println("Starting...")
	start := time.Now()
	waitgroup.Add(2)
	go stingy()
	go spendy()
	waitgroup.Wait()
	println(money)
	fmt.Printf("Time elapsed %v\n", time.Since(start))
}