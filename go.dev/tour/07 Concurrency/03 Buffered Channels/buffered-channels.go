package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 This will throw a panic, since the channel buffer size is 2.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
