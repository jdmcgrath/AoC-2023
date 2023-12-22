package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"strings"
)

func main() {
	processFile("./day9/input.txt")

}
func processFile(path string) {
	readFile, err := os.Open(path)
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	total := 0
	for fileScanner.Scan() {
		total += processLine(fileScanner.Text())
	}
	fmt.Println(total)
}

func processLine(line string) int {
	numberPyramid := make([][]int, 0)
	numbers := strings.Split(line, " ")
	currentLine := toInts(numbers)
	numberPyramid = append(numberPyramid, currentLine)
	processIntSlice(currentLine, &numberPyramid)

	extrapolatedAdditions := extrapolateNumberPyramid(numberPyramid)
	//fmt.Println(extrapolatedAdditions)
	//fmt.Println(extrapolatedAdditions[0][len(extrapolatedAdditions[0])-1])
	return extrapolatedAdditions[0][len(extrapolatedAdditions[0])-1]
}

func toInts(stringSlice []string) (intSlice []int) {
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		aoc.Check(err)
		intSlice = append(intSlice, num)
	}
	return
}

func findDiffBetweenIntsInSlice(slice []int) (diff []int) {
	for i := 0; i < len(slice)-1; i++ {
		firstNumber := slice[i]
		secondNumber := slice[i+1]
		differenceBetween := secondNumber - firstNumber
		diff = append(diff, differenceBetween)
	}
	return
}

func isWholeSliceZeros(slice []int) bool {
	for i := range slice {
		if slice[i] != 0 {
			return false
		}
	}
	return true
}

func processIntSlice(intSlice []int, numberPyramid *[][]int) {
	if !isWholeSliceZeros(intSlice) {
		nextLine := findDiffBetweenIntsInSlice(intSlice)
		*numberPyramid = append(*numberPyramid, nextLine)
		processIntSlice(nextLine, numberPyramid)
	}
}

func extrapolateNumberPyramid(numberPyramid [][]int) [][]int {
	for i := len(numberPyramid) - 2; i >= 0; i-- {
		lastDigitFromOneBeforeThis := numberPyramid[i+1][len(numberPyramid[i+1])-1]
		lastDigitFromThis := numberPyramid[i][len(numberPyramid[i])-1]
		addition := lastDigitFromThis + lastDigitFromOneBeforeThis
		numberPyramid[i] = append(numberPyramid[i], addition)
	}
	return numberPyramid
}
