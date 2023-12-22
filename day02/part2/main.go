package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"strconv"
	"strings"
)

type Pull map[string]int
type Game struct {
	ID    int
	Pulls []Pull
}

func main() {
	file, err := os.Open("./day2/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(file.Close())
	}()

	scanner := bufio.NewScanner(file)

	var total int
	for gameIndex := 1; scanner.Scan(); gameIndex++ {
		bagOfCurrentGame := make(map[string]int)
		line := scanner.Text()
		game := parseToGame(gameIndex, line)

		for _, pull := range game.Pulls {
			bagOfCurrentGame = checkPullAndUpdateBag(pull, bagOfCurrentGame)
		}
		valueOfGame := 1
		for _, value := range bagOfCurrentGame {
			valueOfGame = valueOfGame * value
		}
		total += valueOfGame
	}

	fmt.Println(total)
}

func parseToGame(gameIndex int, line string) Game {
	lineParts := strings.Split(line, ": ")
	pullsData := strings.Split(lineParts[1], "; ")

	var pulls []Pull
	for _, pullData := range pullsData {
		pull := parseToPull(pullData)
		pulls = append(pulls, pull)
	}

	return Game{
		ID:    gameIndex,
		Pulls: pulls,
	}
}

func parseToPull(data string) Pull {
	pull := Pull{}
	items := strings.Split(data, ", ")
	for _, item := range items {
		itemParts := strings.Split(item, " ")
		amount, _ := strconv.Atoi(itemParts[0])
		color := itemParts[1]
		pull[color] = amount
	}
	return pull
}

func checkPullAndUpdateBag(pull Pull, bag map[string]int) map[string]int {
	for color, amount := range pull {
		if amount > bag[color] {
			bag[color] = amount
		}
	}
	return bag
}
