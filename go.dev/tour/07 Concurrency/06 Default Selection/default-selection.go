package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	bomb := time.After(5000 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick")
		case <- bomb:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}