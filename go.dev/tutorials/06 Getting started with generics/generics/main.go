package main

import "fmt"

func sumIntegersOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V = 0

	for _, v := range m {
		sum += v
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

	fmt.Printf(
		"Generic sum results:\nintegers = %v\nfloats = %v",
		sumIntegersOrFloats(integers),
		sumIntegersOrFloats(floats),
	)
}
