package main

import "fmt"

func fibonacci(numberChannel, quitChannel chan int) {
	x, y := 0, 1

	for {
		select {
		case numberChannel <- x:
			x, y = y, x + y
		case <- quitChannel:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	myNumberChannel := make(chan int)
	myQuitChannel := make(chan int)

	go func () {
		for i := 0; i < 10; i += 1 {
			fmt.Println(<- myNumberChannel)
		}

		myQuitChannel <- 0
	}()

	fibonacci(myNumberChannel, myQuitChannel)
}