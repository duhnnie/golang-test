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

var windRegex = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
var tafValidation = regexp.MustCompile(`.*TAF.*`)
var comment = regexp.MustCompile(`\w*#.*`)
var metarClose = regexp.MustCompile(`.*=`)
var variableWind = regexp.MustCompile(`.*VRB\d{2}KT`)
var validWind = regexp.MustCompile(`\d{5}KT`)
var windDirOnly = regexp.MustCompile(`(\d{3})\d{2}KT`)
var windDist [8]int

func parseToArray(text string) []string {
	lines := strings.Split(text, "\n")
	metarSlice := make([]string, len(lines))
	metarStr := ""

	for _, line := range lines {
		if tafValidation.MatchString(line) {
			break
		}

		if !comment.MatchString(line) {
			metarStr += strings.Trim(line, " ")
		}

		if metarClose.MatchString(line) {
			metarSlice = append(metarSlice, metarStr)
			metarStr = ""
		}
	}

	return metarSlice
}

func extractWindDirection(metars []string) []string {
	winds := make([]string, 0, len(metars))

	for _, metar := range metars {
		if windRegex.MatchString(metar) {
			winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
		}
	}

	return winds
}

func mineWindDistribution(winds []string) {
	for _, wind := range winds {
		if variableWind.MatchString(wind) {
			for i := 0; i < 8; i++ {
				windDist[i]++
			}
		} else if validWind.MatchString(wind) {
			windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]

			if d, err := strconv.ParseFloat(windStr, 64); err == nil {
				dirIndex := int(math.Round(d/45.0)) % 8
				windDist[dirIndex] ++
			}
		}
	}
}

func main() {
	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()

	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))

		if err != nil {
			panic(err)
		}

		text := string(dat)

		// 1. Change to array, each metar report is a separate item in the array.
		metarReports := parseToArray(text)

		// 2. Extract wind direction
		windDirections := extractWindDirection(metarReports)

		// 3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
		mineWindDistribution(windDirections)
	}

	elapsed := time.Since(start)
	fmt.Printf("%v\n", windDist)
	fmt.Printf("Processing took %s\n", elapsed)
}