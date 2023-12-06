package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"unicode"
)

type Number struct {
	value        int
	yCoordinate  int
	xCoordinates []int
}

type Symbol struct {
	xCoordinate int
	yCoordinate int
}

func OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	aoc.Check(err)
	return file
}

func ParseSymbolsAndNumbers(scanner *bufio.Scanner) (map[Symbol]bool, []Number) {
	symbols := make(map[Symbol]bool)
	var numbers []Number

	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		numbersOnLine := parseNumbersAndSymbolsOnLine(lineNumber, scanner.Text(), symbols)
		numbers = append(numbers, numbersOnLine...)
	}

	return symbols, numbers
}

func CalculateSum(symbols map[Symbol]bool, numbers []Number) int {
	sum := 0
	for _, num := range numbers {
		if isAdjacentToSymbol(num, symbols) {
			sum += num.value
		}
	}

	return sum
}

func main() {
	file := OpenFile("./day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	symbols, numbers := ParseSymbolsAndNumbers(scanner)

	sum := CalculateSum(symbols, numbers)

	fmt.Println("Sum of all part numbers:", sum)
}

func parseNumbersAndSymbolsOnLine(lineNumber int, line string, symbols map[Symbol]bool) []Number {
	var numbersOnLine []Number
	currentNumber := ""

	for x, char := range line {
		if unicode.IsDigit(char) {
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				num, _ := strconv.Atoi(currentNumber)
				numbersOnLine = append(numbersOnLine, Number{
					value:        num,
					yCoordinate:  lineNumber,
					xCoordinates: []int{x - len(currentNumber), x - 1},
				})
				currentNumber = ""
			}

			if char != '.' {
				symbols[Symbol{xCoordinate: x, yCoordinate: lineNumber}] = true
			}
		}
	}

	if currentNumber != "" {
		num, err := strconv.Atoi(currentNumber)
		aoc.Check(err)
		startCoordinate := len(line) - len(currentNumber)
		numbersOnLine = append(numbersOnLine, Number{
			value:        num,
			yCoordinate:  lineNumber,
			xCoordinates: []int{startCoordinate, len(line) - 1},
		})
	}

	return numbersOnLine
}

func isAdjacentToSymbol(num Number, symbols map[Symbol]bool) bool {
	for _, x := range num.xCoordinates {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if _, exists := symbols[Symbol{xCoordinate: x + dx, yCoordinate: num.yCoordinate + dy}]; exists {
					return true
				}
			}
		}
	}

	return false
}
