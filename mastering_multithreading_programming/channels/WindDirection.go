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

func parseToArray(inputChannel chan string, outputChannel chan []string) {
	for text := range inputChannel {
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

		outputChannel <- metarSlice
	}

	close(outputChannel)
}

func extractWindDirection(inputChannel chan []string, outputChannel chan []string) {
	for metars := range inputChannel {
		winds := make([]string, 0, len(metars))

		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}

		outputChannel <- winds
	}

	close(outputChannel)
}

func mineWindDistribution(inputChannel chan []string, outputChannel chan [8]int) {
	for winds := range inputChannel {
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

	outputChannel <- windDist
}

func main() {
	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()

	textChannel := make(chan string)
	metarsChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultChannel := make(chan [8]int)

	go parseToArray(textChannel, metarsChannel)
	go extractWindDirection(metarsChannel, windsChannel)
	go mineWindDistribution(windsChannel, resultChannel)

	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))

		if err != nil {
			panic(err)
		}

		textChannel <- string(dat)
	}

	close(textChannel)

	elapsed := time.Since(start)
	fmt.Printf("%v\n", <- resultChannel)
	fmt.Printf("Processing took %s\n", elapsed)
}