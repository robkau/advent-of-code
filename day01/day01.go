package main

import (
	"fmt"
	"github.com/robkau/advent-of-code/util"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// --- Day 1: Chronal Calibration ---
// https://adventofcode.com/2018/day/1

const (
	InputDataUrl = "https://adventofcode.com/2018/day/1/input"
)

func applyFrequencyChanges(baseFrequency int, frequencyChanges []int) int {
	for _, fc := range frequencyChanges {
		baseFrequency += fc
	}
	return baseFrequency
}

func findFirstFrequencyCollision(baseFrequency int, frequencyChanges []int) int {
	var Present struct{}
	visitedFrequencies := make(map[int]struct{})

	for {
		for _, fc := range frequencyChanges {
			// loop forever until the first repeat is found
			_, found := visitedFrequencies[baseFrequency]
			if found {
				return baseFrequency
			}
			visitedFrequencies[baseFrequency] = Present
			baseFrequency += fc
		}
	}
}

func main() {
	// get input bytes
	inputBytes, err := util.HttpRequestWithCookie(InputDataUrl, http.MethodGet, util.SessionCookie, os.Getenv(util.SessionEnvVar))
	if err != nil {
		panic(err)
	}
	// convert input bytes to lines of strings, cut off trailing newline and then split by newline
	inputs := strings.Split(strings.TrimSuffix(string(inputBytes), util.LineBreak), util.LineBreak)

	// convert input strings to integer data
	frequencyChanges := make([]int, 0)
	for _, s := range inputs {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Failed to convert input string to int")
			panic(err)
		}
		frequencyChanges = append(frequencyChanges, i)
	}

	//------------------
	// Question 1:
	// Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

	// Strategy:
	// initialize frequency and then then add or subtract each frequency change
	var baseFrequency int = 0
	finalFrequency := applyFrequencyChanges(baseFrequency, frequencyChanges)

	// Answer:
	println("Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied? : ", finalFrequency)

	//------------------
	// Question 2:
	// What is the first frequency your device reaches twice, while repeatedly looping over the input frequency changes?

	// Strategy:
	// Keep a map in memory to store membership of visited frequencies, record the first collision
	// If this grows too large look into a filter data structure
	firstCollisionFrequency := findFirstFrequencyCollision(baseFrequency, frequencyChanges)

	// Answer:
	println("What is the first frequency your device reaches twice? : ", firstCollisionFrequency)
}
