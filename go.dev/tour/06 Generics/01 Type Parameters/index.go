package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](haystack []T, needle T) int {
	for i := 0; i < len(haystack); i ++ {
		if haystack[i] == needle {
			return i
		}
	}

	return -1
}

func main() {
	numbers := []int{23, 76, -23, 1}
	fmt.Println(Index(numbers, -23))

	words := []string{"I", "Love", "my", "mom!"}
	fmt.Println(Index(words, "cat"))
	fmt.Println(Index(words, "mom!"))
}