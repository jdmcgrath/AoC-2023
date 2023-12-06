package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	aoc "github.com/jdmcgrath/AoC-2023"
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

type NumberID struct {
	value       int
	coordinates [2]int // x and y coordinates
}

type FileHandler struct{}

type SymbolNumberParser struct{}

type Calculator struct {
	symbols map[Symbol]bool
	numbers []Number
}

func NewCalculator(symbols map[Symbol]bool, numbers []Number) *Calculator {
	return &Calculator{symbols: symbols, numbers: numbers}
}

func (fh *FileHandler) OpenFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	aoc.Check(err)
	return file
}

func (sp *SymbolNumberParser) ParseSymbolsAndNumbers(scanner *bufio.Scanner) (map[Symbol]bool, []Number) {
	symbols := make(map[Symbol]bool)
	var numbers []Number

	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		parseNumbersAndSymbolsOnLine := sp.parseNumbersAndSymbolsOnLine(lineNumber, scanner.Text(), symbols)
		numbers = append(numbers, parseNumbersAndSymbolsOnLine...)
	}

	return symbols, numbers
}

func (sp *SymbolNumberParser) parseNumbersAndSymbolsOnLine(lineNumber int, line string, symbols map[Symbol]bool) []Number {
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

			if char == '*' {
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

func (c *Calculator) CalculateSum() int {
	sum := 0
	numsFromSymbol := make(map[Symbol][]NumberID)
	for _, num := range c.numbers {
		c.getNumsBySymbol(num, numsFromSymbol)
	}

	for _, numberIDs := range numsFromSymbol {
		if len(numberIDs) != 2 {
			continue
		}
		multiplication := 1
		for _, id := range numberIDs {
			multiplication *= id.value
		}
		sum += multiplication
	}

	return sum
}

func (c *Calculator) getNumsBySymbol(num Number, numsFromSymbol map[Symbol][]NumberID) {
	for _, x := range num.xCoordinates {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				symbol := Symbol{xCoordinate: x + dx, yCoordinate: num.yCoordinate + dy}
				if _, exists := c.symbols[symbol]; exists {
					numberID := NumberID{value: num.value, coordinates: [2]int{num.xCoordinates[0], num.yCoordinate}}
					if !c.containsNumberID(numsFromSymbol[symbol], numberID) {
						numsFromSymbol[symbol] = append(numsFromSymbol[symbol], numberID)
					}
				}
			}
		}
	}
}

func (c *Calculator) containsNumberID(slice []NumberID, id NumberID) bool {
	for _, item := range slice {
		if item == id {
			return true
		}
	}
	return false
}

func main() {
	fh := FileHandler{}
	file := fh.OpenFile("./day3/input.txt")
	defer func() {
		err := file.Close()
		aoc.Check(err)
	}()

	scanner := bufio.NewScanner(file)
	sp := SymbolNumberParser{}
	symbols, numbers := sp.ParseSymbolsAndNumbers(scanner)

	calc := NewCalculator(symbols, numbers)
	sum := calc.CalculateSum()

	fmt.Println("Sum of all part numbers:", sum)
}
