package main

import (
	"fmt"
	"time"
)

var money = 100

func stingy() {
	for i := 1; i <= 1000; i++ {
		money = money + 10
		time.Sleep(1 * time.Millisecond)
	}

	println("Stingy done!")
}

func spendy() {
	for i:= 1; i <= 1000; i++ {
		money = money - 10
		time.Sleep(1 * time.Millisecond)
	}

	println("Spendy done!")
}

func main() {
	fmt.Println("Starting...")
	go stingy()
	go spendy()
	time.Sleep(30000 * time.Millisecond)
	println(money)
}