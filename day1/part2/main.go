package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	processFile("./day1/input.txt")
}

func processFile(path string) {
	readFile, err := os.Open(path)
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	runningCount := 0
	for fileScanner.Scan() {
		runningCount += processLine(fileScanner.Text())
	}
	fmt.Println(runningCount)
}

func processLine(line string) int {
	digitsInString := getDigitsInString(line)
	digitsString := strings.Join(digitsInString, "")
	first, last := getFirstAndLastFromString(digitsString)
	justTwoDigits := first + last
	calib, err := strconv.Atoi(justTwoDigits)
	aoc.Check(err)
	println(calib)
	return calib
}

func getDigitsInString(s string) []string {
	var digits []string
	var wordBuilder strings.Builder
	textNumbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, string(r))
			wordBuilder.Reset()
		} else if unicode.IsLetter(r) {
			wordBuilder.WriteRune(r)
			// Check if wordBuilder contains any of the text numbers
			for word, number := range textNumbers {
				if strings.Contains(wordBuilder.String(), word) {
					digits = append(digits, number)
					wordBuilder.Reset()
					break
				}
			}
		} else {
			wordBuilder.Reset()
		}
	}

	// Check for the last word after the loop
	for word, number := range textNumbers {
		if strings.Contains(wordBuilder.String(), word) {
			digits = append(digits, number)
			break
		}
	}

	return digits
}

func getFirstAndLastFromString(inputString string) (first, last string) {
	if len(inputString) > 0 {
		first = string(inputString[0])
		last = string(inputString[len(inputString)-1])
	}
	return first, last
}
