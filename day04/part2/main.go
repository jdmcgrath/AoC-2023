package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"strings"
)

type scratchie struct {
	id      int
	numbers string
}

func main() {
	file, err := os.Open("./day4/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(file.Close())
	}()

	scanner := bufio.NewScanner(file)

	var total int
	scratchieStack := make(map[int]int)
	for stratchieId := 1; scanner.Scan(); stratchieId++ {
		scratchieStack[stratchieId] += 1
		fmt.Println("number of this card: ", scratchieStack[stratchieId])
		numbers := strings.Split(scanner.Text(), ": ")[1]
		winningAndHave := strings.Split(numbers, " | ")
		winningNumbersAsString := winningAndHave[0]
		numbersHaveAsString := winningAndHave[1]

		winningNumbers := splitNumbersSeparatedBySpaces(winningNumbersAsString)
		numbersHave := splitNumbersSeparatedBySpaces(numbersHaveAsString)
		numberOfWins := countWinningNumbers(winningNumbers, numbersHave)
		for i := numberOfWins; i > 0; i-- {
			scratchieStack[stratchieId+i] += scratchieStack[stratchieId]
		}
	}
	for _, v := range scratchieStack {
		total += v
	}
	fmt.Println(total)
}

func splitNumbersSeparatedBySpaces(numbers string) []int {
	var numbersAsInt []int
	numbersAsStrings := strings.Split(numbers, " ")
	for _, numberAsString := range numbersAsStrings {
		if numberAsString != "" {
			num, err := strconv.Atoi(numberAsString)
			aoc.Check(err)
			numbersAsInt = append(numbersAsInt, num)
		}
	}
	return numbersAsInt
}

func countWinningNumbers(winningNumbers, numbersHave []int) int {
	winningNumbersMap := make(map[int]bool)
	for _, winningNumber := range winningNumbers {
		winningNumbersMap[winningNumber] = true
	}
	count := 0
	for _, number := range numbersHave {
		if winningNumbersMap[number] {
			count++
		}
	}
	return count
}
