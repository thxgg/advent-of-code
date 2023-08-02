package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func score(s string) int {
	sn := s[0]
	if sn >= 'a' && sn <= 'z' {
		return int(sn - 'a' + 1)
	} else if sn >= 'A' && sn <= 'Z' {
		return int(sn - 'A' + 27)
	} else {
		return 0
	}
}

func partOne() {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		seenFirst, seenSecond := make(map[string]bool), make(map[string]bool)
		split := strings.Split(line, "")
		first, second := split[:len(split)/2], split[len(split)/2:]

		for i := 0; i < len(split)/2; i++ {
			if seenSecond[first[i]] {
				sum += score(first[i])
				break
			} else {
				seenFirst[first[i]] = true
			}
			if seenFirst[second[i]] {
				sum += score(second[i])
				break
			} else {
				seenSecond[second[i]] = true
			}
		}
	}

	fmt.Println(sum)
}

func partTwo() {
	sum := 0

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		seen := make(map[string]byte)
		first, second, third := lines[i], lines[i+1], lines[i+2]
		for _, letter := range strings.Split(first, "") {
			seen[letter] |= 0b001
			if seen[letter] == 0b111 {
				sum += score(letter)
				break
			}
		}
		for _, letter := range strings.Split(second, "") {
			seen[letter] |= 0b010
			if seen[letter] == 0b111 {
				sum += score(letter)
				break
			}
		}
		for _, letter := range strings.Split(third, "") {
			seen[letter] |= 0b100
			if seen[letter] == 0b111 {
				sum += score(letter)
				break
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	//partOne()
	partTwo()
}
