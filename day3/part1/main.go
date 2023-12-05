package main

import (
	"bufio"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"unicode"
)

type number struct {
	coordinateX, coordinateY, number int
}

func main() {
	file, err := os.Open("./day3/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(file.Close())
	}()
	var fileOfChars []string
	scanner := bufio.NewScanner(file)
	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		fileOfChars = append(fileOfChars, scanner.Text())
	}
}

func parseNumbersAndCoordinatesOnLine(line []rune) []number {
	inNumber := false
	var number []rune
	for xCoordinate, character := range line {
		if isDigit(character) {
			number = append(number, character)
		}
	}
}

func isDigit(character rune) bool {
	return unicode.IsDigit(character)
}
