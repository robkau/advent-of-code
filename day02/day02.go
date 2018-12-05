package main

import (
	"fmt"
	"github.com/robkau/advent-of-code/util"
	"net/http"
	"os"
	"strings"
)

// --- Day 2: Inventory Management System ---
// https://adventofcode.com/2018/day/2

const (
	InputDataUrl = "https://adventofcode.com/2018/day/2/input"
)

func numRepeatingNLetters(inputs []string, n int) (total int) {
	// returns how many items from 'inputs' contain exactly n occurrences of any character
	for _, i := range inputs {
		if containsNOccurrences(i, n) {
			total += 1
		}
	}
	return total
}

func containsNOccurrences(input string, n int) bool {
	// returns true if input contains exactly n occurrences of any character
	seen := make(map[string]int)
	for _, c := range input {
		seen[string(c)] += 1
	}

	for _, c := range input {
		if seen[string(c)] == n {
			return true
		}
	}
	return false
}

func roughlyEqual(a string, b string, allowedErrors int) bool {
	// returns true if a and b differ by , at most, allowedErrors characters
	if len(a) != len(b) {
		return false
	}
	errors := 0
	for i := range a {
		if a[i] != b[i] {
			errors += 1
			if errors > allowedErrors {
				return false
			}
		}
	}
	return true
}

func matchingCharacters(a string, b string) (matching string) {
	for i := range a {
		if a[i] == b[i] {
			matching += string(a[i])
		}
	}
	return matching
}

func main() {
	// get input bytes
	inputBytes, err := util.HttpRequestWithCookie(InputDataUrl, http.MethodGet, util.SessionCookie, os.Getenv(util.SessionEnvVar))
	if err != nil {
		panic(err)
	}
	// convert input bytes to usable strings: cut off trailing newline and then split by newline
	inputs := strings.Split(strings.TrimSuffix(string(inputBytes), util.LineBreak), util.LineBreak)

	//------------------
	// Question 1:
	// What is the checksum for your list of box IDs?

	// Strategy:
	// find how many box IDs have two repeating chars
	// find how many box IDs have three repeating chars
	// multiply the result

	// Answer:
	twoRepeats := numRepeatingNLetters(inputs, 2)
	threeRepeats := numRepeatingNLetters(inputs, 3)
	fmt.Println("Checksum: ", twoRepeats*threeRepeats)

	//------------------
	// Question 2:
	// you're ready to find the boxes full of prototype fabric.
	// the boxes will have IDs which differ by exactly one character at the same position in both strings.
	// what letters are common between the two correct box IDs?

	// Strategy:
	// It's only 250 short strings, compare each to others until a match is found

	// Answer:
	for i, si := range inputs {
		for j, sj := range inputs {
			if i == j {
				continue
			}
			if roughlyEqual(si, sj, 1) {
				fmt.Println("Matching characters: ", matchingCharacters(si, sj))
				return
			}
		}
	}
	fmt.Println("Finished with no answer")
}
