package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne() {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		ranges := strings.Split(line, ",")
		first, second := ranges[0], ranges[1]
		firstEnds, secondEnds := strings.Split(first, "-"), strings.Split(second, "-")
		firstStart, _ := strconv.Atoi(firstEnds[0])
		firstEnd, _ := strconv.Atoi(firstEnds[1])
		secondStart, _ := strconv.Atoi(secondEnds[0])
		secondEnd, _ := strconv.Atoi(secondEnds[1])
		if firstStart <= secondStart && firstEnd >= secondEnd {
			count++
		} else if secondStart <= firstStart && secondEnd >= firstEnd {
			count++
		}
	}

	fmt.Println(count)
}

func partTwo() {
	count := 0

	for _, line := range strings.Split(input, "\n") {
		ranges := strings.Split(line, ",")
		first, second := ranges[0], ranges[1]
		firstEnds, secondEnds := strings.Split(first, "-"), strings.Split(second, "-")
		firstStart, _ := strconv.Atoi(firstEnds[0])
		firstEnd, _ := strconv.Atoi(firstEnds[1])
		secondStart, _ := strconv.Atoi(secondEnds[0])
		secondEnd, _ := strconv.Atoi(secondEnds[1])
		if firstStart <= secondStart && secondStart <= firstEnd {
			count++
		} else if secondStart <= firstStart && firstStart <= secondEnd {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	//partOne()
	partTwo()
}
