package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countFrecuency(url string, frecuency *[26]int) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	for _, b := range body {
		letter := strings.ToLower(string(b))
		index := strings.Index(allLetters, letter)

		if index >= 0 {
			frecuency[index] += 1
		}
	}
}

func main() {
	var frecuency [26]int
	start := time.Now()

	for i := 1000; i <= 1200; i += 1 {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%v.txt", i)
		countFrecuency(url, &frecuency)
	}

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)

	for index, letter := range allLetters {
		fmt.Printf("%v -> %v\n", string(letter), frecuency[index])
	}
}