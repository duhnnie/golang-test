package main

import (
	"fmt"
	"time"
)

func fibonacci(number int, channel chan int) {
	x, y := 0, 1

	for i := 0; i < number; i += 1 {
		channel <- x
		x, y = y, x + y
		time.Sleep(100 * time.Millisecond)
	}

	close(channel)
}

func main() {
	myChannel := make(chan int, 10)

	go fibonacci(cap(myChannel), myChannel)

	for v := range myChannel {
		fmt.Println(v)
	}
}