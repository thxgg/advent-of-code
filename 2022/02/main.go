package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func scoreRound(round string) int {
	choices := strings.Split(round, " ")
	opp, mine := choices[0], choices[1]

	score := 0

	switch mine {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	switch mine {
	case "X":
		switch opp {
		case "A":
			score += 3
		case "B":
			score += 0
		case "C":
			score += 6
		}
	case "Y":
		switch opp {
		case "A":
			score += 6
		case "B":
			score += 3
		case "C":
			score += 0
		}
	case "Z":
		switch opp {
		case "A":
			score += 0
		case "B":
			score += 6
		case "C":
			score += 3
		}
	}

	return score
}

func partOne() {
	score := 0

	for _, round := range strings.Split(input, "\n") {
		if len(round) > 0 {
			score += scoreRound(round)
		}
	}

	fmt.Println(score)
}

func scoreRound2(round string) int {
	choices := strings.Split(round, " ")
	opp, outcome := choices[0], choices[1]

	switch opp {
	case "A":
		switch outcome {
		case "X":
			return 0 + 3
		case "Y":
			return 3 + 1
		case "Z":
			return 6 + 2
		}
	case "B":
		switch outcome {
		case "X":
			return 0 + 1
		case "Y":
			return 3 + 2
		case "Z":
			return 6 + 3
		}
	case "C":
		switch outcome {
		case "X":
			return 0 + 2
		case "Y":
			return 3 + 3
		case "Z":
			return 6 + 1
		}
	}

	return 0
}

func partTwo() {
	score := 0

	for _, round := range strings.Split(input, "\n") {
		if len(round) > 0 {
			score += scoreRound2(round)
		}
	}

	fmt.Println(score)
}

func main() {
	// partOne()
	partTwo()
}
