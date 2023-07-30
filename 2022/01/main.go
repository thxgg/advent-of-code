package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne() {
	max := 0

	for _, elf := range strings.Split(input, "\n\n") {
		calories := 0

		for _, food := range strings.Split(elf, "\n") {
			n, _ := strconv.Atoi(food)
			calories += n
		}

		if calories > max {
			max = calories
		}
	}

	fmt.Println(max)
}

func partTwo() {
	var elves []int

	for _, elf := range strings.Split(input, "\n\n") {
		calories := 0

		for _, food := range strings.Split(elf, "\n") {
			n, _ := strconv.Atoi(food)
			calories += n
		}

		elves = append(elves, calories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	sum := 0
	for _, val := range elves[:3] {
		sum += val
	}
	fmt.Println(sum)
}

func main() {
	// partOne()
	partTwo()
}
