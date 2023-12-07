package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2023"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Hand represents a hand of cards along with its bid
type Hand struct {
	Cards string
	Bid   int
	Rank  int
	Type  HandType
}

// HandType represents the type of hand in terms of strength
type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// parseInput reads the input file and returns a slice of Hands
func parseInput(filePath string) ([]Hand, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		aoc.Check(file.Close())
	}()
	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // Skip lines with invalid bids
		}
		hands = append(hands, Hand{Cards: parts[0], Bid: bid})
	}
	return hands, scanner.Err()
}

// classifyAndSortHands classifies the hands based on their type and sorts them
func classifyAndSortHands(hands []Hand) {
	for i, hand := range hands {
		hands[i].Type = classifyHand(hand.Cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type != hands[j].Type {
			return hands[i].Type > hands[j].Type
		}
		return compareHands(hands[i].Cards, hands[j].Cards) > 0
	})

	// Assign ranks based on sorted order
	for i := range hands {
		hands[i].Rank = len(hands) - i
	}
}

// classifyHand classifies the hand and returns its type
func classifyHand(cards string) HandType {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card]++
	}

	switch len(counts) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, count := range counts {
			if count == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for _, count := range counts {
			if count == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

// compareHands compares two hands based on their card strengths
func compareHands(hand1, hand2 string) int {
	for i := 0; i < len(hand1); i++ {
		if hand1[i] != hand2[i] {
			return cardStrength(rune(hand1[i])) - cardStrength(rune(hand2[i]))
		}
	}
	return 0
}

// cardStrength returns the strength of a card
func cardStrength(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

// calculateTotalWinnings calculates the total winnings
func calculateTotalWinnings(hands []Hand) int {
	total := 0
	for _, hand := range hands {
		total += hand.Bid * hand.Rank
	}
	return total
}

func main() {
	hands, err := parseInput("./day7/input.txt")
	aoc.Check(err)
	classifyAndSortHands(hands)
	totalWinnings := calculateTotalWinnings(hands)
	fmt.Println("Total winnings:", totalWinnings)
}
