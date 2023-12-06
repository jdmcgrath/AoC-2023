package main

import "fmt"

type race struct {
	maxTime        int
	recordDistance int
}

func main() {
	maxTime := 45977295
	recordDistance := 305106211101695
	waysToWin := totalWaysToWin(findFirstWin(maxTime, recordDistance), findLastWin(maxTime, recordDistance))
	fmt.Println(waysToWin)
}

// Binary search could make this more efficient
func findFirstWin(timeGiven, distanceRecord int) (firstWin int) {
	foundWin := false
	// msHeld is interchangeable to speedOfTravel
	for msHeld := 1; foundWin == false; msHeld++ {
		if doesBeatRecordDistance(distanceRecord, totalDistanceTravelled(timeGiven-msHeld, msHeld)) {
			foundWin = true
			firstWin = msHeld
		}
	}
	return firstWin
}

// Binary search could make this more efficient
func findLastWin(timeGiven, distanceRecord int) (lastWin int) {
	foundWin := false
	// msHeld is interchangeable to speedOfTravel
	for msHeld := timeGiven - 1; foundWin == false; msHeld-- {
		if doesBeatRecordDistance(distanceRecord, totalDistanceTravelled(timeGiven-msHeld, msHeld)) {
			foundWin = true
			lastWin = msHeld
		}
	}
	return lastWin
}

func doesBeatRecordDistance(record, actual int) bool {
	if record < actual {
		return true
	} else {
		return false
	}
}

func totalDistanceTravelled(remainingTime int, speedOfTravel int) (distance int) {
	return remainingTime * speedOfTravel
}

func totalWaysToWin(lowestSuccessfulHoldMs, highestSuccessfulHoldMs int) int {
	return highestSuccessfulHoldMs - lowestSuccessfulHoldMs + 1
}
