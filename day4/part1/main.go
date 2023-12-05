package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"math"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("./day4/input.txt")
	aoc.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int64
	var totalMutex sync.Mutex
	var workerGroup sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()

		// Launch a goroutine to process each line.
		workerGroup.Add(1)
		go func(str string) {
			defer workerGroup.Done()

			parts := strings.Split(str, ": ")[1]
			winningNumbersAsString, numbersHaveAsString := splitParts(parts)

			winningNumbers := make(map[string]bool)
			for _, winningNumber := range strings.Fields(winningNumbersAsString) {
				winningNumbers[winningNumber] = true
			}

			numberOfWinningNumbers := countWinningNumbers(winningNumbers, strings.Fields(numbersHaveAsString))

			pointsOfScratchie := calculatePoints(numberOfWinningNumbers)

			totalMutex.Lock()
			total += pointsOfScratchie
			totalMutex.Unlock()
		}(line)
	}

	workerGroup.Wait()

	fmt.Println(total)
}

func splitParts(parts string) (winningNumbersAsString, numbersHaveAsString string) {
	delimeterIndex := strings.Index(parts, " | ")
	winningNumbersAsString = parts[:delimeterIndex]
	numbersHaveAsString = parts[delimeterIndex+3:]
	return
}

func countWinningNumbers(winningNumbers map[string]bool, numbersHave []string) (count int) {
	for _, number := range numbersHave {
		if winningNumbers[number] {
			count++
		}
	}
	return
}

func calculatePoints(winningNumbersCount int) int64 {
	if winningNumbersCount == 0 {
		return 0
	}
	return int64(math.Pow(2, float64(winningNumbersCount-1)))
}
