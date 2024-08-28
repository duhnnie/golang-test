package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

var r = regexp.MustCompile(`\((\d*),(\d*)\)`)
var numberOfThreads = 8
var waitGroup = sync.WaitGroup{}

func findArea(inputChannel chan string) {
	for pointsStr := range inputChannel {
		var points []Point2D
		area := 0.0

		for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{ x, y })
		}

		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i + 1) % len(points)]
			area += float64(a.x * b.y) - float64(a.y * b.x)
		}

		fmt.Println(math.Abs(area) / 2.0)
	}

	waitGroup.Done()
}

func main() {
	start := time.Now()
	filePath, _ := filepath.Abs("./polygons.txt")
	file, _ := ioutil.ReadFile(filePath)
	text := string(file)

	inputChannel := make(chan string, 1000)

	for i := 0; i < numberOfThreads; i ++ {
		waitGroup.Add(1)
		go findArea(inputChannel)
	}

	for _, line := range strings.Split(text , "\n") {
		inputChannel <- line
	}

	close(inputChannel)
	waitGroup.Wait()
	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", timeElapsed)
}