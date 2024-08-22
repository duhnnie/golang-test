package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var matches []string
var waitGroup = sync.WaitGroup{}
var lock = sync.Mutex{}

func searchFile(directory, filename string) {
	fmt.Println("Searching in", directory)

	files, _ := ioutil.ReadDir(directory)

	for _, file := range files {
		path := filepath.Join(directory, file.Name())

		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, path)
			lock.Unlock()
		} else if file.IsDir() {
			waitGroup.Add(1)
			go searchFile(path, filename)
		}
	}

	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)
	go searchFile("/Users/daniel/Projects", "README.md")
	waitGroup.Wait()

	for _, match := range matches {
		fmt.Printf("Match: %v\n", match)
	}

	fmt.Printf("%v matches found.", len(matches))
}