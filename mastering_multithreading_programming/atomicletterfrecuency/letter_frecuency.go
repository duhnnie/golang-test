package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"
var lock = sync.Mutex{}
var waitGroup = sync.WaitGroup{}
var threads = 8

func countFrecuency(url string, frecuency *[26]int64) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	for i := 0; i < 20; i++ {
		for _, b := range body {
			letter := strings.ToLower(string(b))
			index := strings.Index(allLetters, letter)

			if index >= 0 {
				// lock.Lock()
				atomic.AddInt64(&frecuency[index], 1)
				// frecuency[index] += 1
				// lock.Unlock()
			}
		}
	}

	waitGroup.Done()
}

func main() {
	var frecuency [26]int64
	start := time.Now()

	for i := 1000; i <= 1200; i += 1 {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%v.txt", i)
		waitGroup.Add(1)
		go countFrecuency(url, &frecuency)
	}

	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)

	for index, letter := range allLetters {
		fmt.Printf("%v -> %v\n", string(letter), frecuency[index])
	}
}