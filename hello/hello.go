package main

import (
	"fmt"
	"example.com/greetings"
	"example.com/math"
	"log"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	sumResult := math.Sum(1, 2)
	fmt.Printf("1 + 2 = %v\n", sumResult)

	multiplyResult := math.Multiply(2, 3)
	fmt.Printf("2 x 3 = %v\n", multiplyResult)
}
