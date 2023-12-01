package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile("[0-9]+")
	runningCount := 0
	for fileScanner.Scan() {
		runningCount += processLine(fileScanner.Text(), re)
	}
	fmt.Println(runningCount)
}

func processLine(line string, re *regexp.Regexp) int {
	digitsInString := re.FindAllString(line, -1)
	digitsString := strings.Join(digitsInString, "")
	first, last := getFirstAndLastFromString(digitsString)
	calibrationString := first + last
	calibrationInt, err := strconv.Atoi(calibrationString)
	aoc.Check(err)
	return calibrationInt
}

func getFirstAndLastFromString(inputString string) (first, last string) {
	if len(inputString) > 0 {
		first = string(inputString[0])
		last = string(inputString[len(inputString)-1])
	}
	return first, last
}
