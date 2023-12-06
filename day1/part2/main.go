package main

import (
	"bufio"
	"errors"
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

func processLine(l string) int {
	textNumbers := map[string]rune{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	firstDigit, err := getFirstDigitInLine(l, textNumbers)
	aoc.Check(err)

	lastDigit, err := getLastDigitInLine(l, textNumbers)
	aoc.Check(err)

	stringNumber := string(firstDigit) + string(lastDigit)
	intNumber, err := strconv.Atoi(stringNumber)
	aoc.Check(err)

	return intNumber
}

func getFirstDigitInLine(l string, digitStrings map[string]rune) (rune, error) {
	var sb strings.Builder
	for _, char := range l {
		if unicode.IsDigit(char) {
			return char, nil
		} else {
			sb.WriteRune(char)
			for digitString, digit := range digitStrings {
				if strings.HasSuffix(sb.String(), digitString) {
					return digit, nil
				}
			}
		}
	}
	return '0', errors.New("could not find a digit")
}

func getLastDigitInLine(l string, digitStrings map[string]rune) (rune, error) {
	var sb strings.Builder
	reversed := reverse(l)
	for _, char := range reversed {
		if unicode.IsDigit(char) {
			return char, nil
		} else {
			sb.WriteRune(char)
			reversedSb := reverse(sb.String())
			for digitString, digit := range digitStrings {
				if strings.HasPrefix(reversedSb, digitString) {
					return digit, nil
				}
			}
		}
	}
	return '0', errors.New("could not find a digit")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
