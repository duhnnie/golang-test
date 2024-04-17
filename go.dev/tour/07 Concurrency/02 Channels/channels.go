package main

import "fmt"

func sum(numbers []int, channel chan int) {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	channel <- sum
}

func main() {
	myNumbers := []int{ 34, -5, 23, 12, 54, 31}

	myChannel := make(chan int)
	go sum(myNumbers[:len(myNumbers)/2], myChannel)
	go sum(myNumbers[len(myNumbers)/2:], myChannel)

	x, y := <- myChannel, <- myChannel

	fmt.Printf("%v + %v = %v\n", x, y, x+y)
}