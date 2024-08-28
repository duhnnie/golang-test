package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Point2D struct {
	x int
	y int
}

var r = regexp.MustCompile(`\((\d*),(\d*)\)`)

func findArea(pointsStr string) float64 {
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


	return math.Abs(area) / 2.0
}

func main() {
	start := time.Now()
	filePath, _ := filepath.Abs("./polygons.txt")
	file, _ := ioutil.ReadFile(filePath)
	text := string(file)

	for _, line := range strings.Split(text , "\n") {
		res := findArea(line)
		fmt.Println(res)
	}

	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", timeElapsed)
}