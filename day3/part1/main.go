package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"unicode"
)

type number struct {
	value        int
	yCoordinate  int
	xCoordinates []int
}

type symbol struct {
	xCoordinate int
	yCoordinate int
}

func main() {
	file, err := os.Open("./day3/input.txt")
	aoc.Check(err)
	defer func() {
		err := file.Close()
		aoc.Check(err)
	}()

	scanner := bufio.NewScanner(file)
	symbols := make(map[symbol]bool)
	var numbers []number

	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		numbersOnLine := parseNumbersAndSymbolsOnLine(lineNumber, scanner.Text(), symbols)
		numbers = append(numbers, numbersOnLine...)
	}

	sum := 0
	for _, num := range numbers {
		if isAdjacentToSymbol(num, symbols) {
			sum += num.value
		}
	}

	fmt.Println("Sum of all part numbers:", sum)
}

func parseNumbersAndSymbolsOnLine(lineNumber int, line string, symbols map[symbol]bool) (numbersOnLine []number) {
	currentNumber := ""
	for x, char := range line {
		if unicode.IsDigit(char) {
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				numbersOnLine = append(numbersOnLine, number{
					value:        num,
					yCoordinate:  lineNumber,
					xCoordinates: []int{x - len(currentNumber), x - 1},
				})
				currentNumber = ""
			}
			if char != '.' {
				symbols[symbol{xCoordinate: x, yCoordinate: lineNumber}] = true
			}
		}
	}
	if currentNumber != "" {
		num, err := strconv.Atoi(currentNumber)
		aoc.Check(err)
		startCoordinate := len(line) - len(currentNumber)
		numbersOnLine = append(numbersOnLine, number{
			value:        num,
			yCoordinate:  lineNumber,
			xCoordinates: []int{startCoordinate, len(line) - 1},
		})
	}

	return numbersOnLine
}

func isAdjacentToSymbol(num number, symbols map[symbol]bool) bool {
	for _, x := range num.xCoordinates {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if _, exists := symbols[symbol{xCoordinate: x + dx, yCoordinate: num.yCoordinate + dy}]; exists {
					return true
				}
			}
		}
	}
	return false
}
