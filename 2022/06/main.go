package main

import (
	_ "embed"
	"fmt"
	_ "github.com/emirpasic/gods/stacks/arraystack"
)

//go:embed input.txt
var input string

func partOne() {
	for i := 0; i < len(input)-3; i++ {
		win := input[i : i+4]
		unique := true
		for j := 0; unique && j < len(win); j++ {
			for k := j + 1; unique && k < len(win); k++ {
				if win[j] == win[k] {
					unique = false
				}
			}
		}
		if unique {
			fmt.Println(i + 4)
			break
		}
	}
}

func partTwo() {
	for i := 0; i < len(input)-13; i++ {
		win := input[i : i+14]
		unique := true
		for j := 0; unique && j < len(win); j++ {
			for k := j + 1; unique && k < len(win); k++ {
				if win[j] == win[k] {
					unique = false
				}
			}
		}
		if unique {
			fmt.Println(i + 14)
			break
		}
	}
}

func main() {
	//partOne()
	partTwo()
}
