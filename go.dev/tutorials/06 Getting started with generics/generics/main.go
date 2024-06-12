package main

import "fmt"

func sumIntegers(numbers map[string]int64) int64 {
	var sum int64 = 0

	for _, i := range numbers {
		sum += i
	}

	return sum
}

func sumFloats(floats map[string]float64) float64 {
	var sum float64 = 0

	for _, i := range floats {
		sum += i
	}

	return sum
}

func main() {
	var integers = map[string]int64{
		"first":  23,
		"second": 45,
	}

	var floats = map[string]float64{
		"first":  45.2,
		"second": 60.1,
	}

	fmt.Printf("sumInts = %v\nsumFloats = %v\n", sumIntegers(integers), sumFloats(floats))
}
