package main

import "fmt"

type race struct {
	maxTime        int
	recordDistance int
}

func main() {
	// For efficiency, we know that in a certain threshold you will always beat the record.
	// To find the threshold, let's find the minimum time you must the car for, and the maximum
	total := 1
	races := []race{
		{
			maxTime:        45,
			recordDistance: 305,
		},
		{
			maxTime:        97,
			recordDistance: 1062,
		}, {
			maxTime:        72,
			recordDistance: 1110,
		}, {
			maxTime:        95,
			recordDistance: 1695,
		}}
	for _, r := range races {
		total = total * totalWaysToWin(findFirstWin(r.maxTime, r.recordDistance), findLastWin(r.maxTime, r.recordDistance))
	}
	fmt.Println(total)
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
