package main

import (
	_ "embed"
	"fmt"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne() {
	notableCycles := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	cycle := 1
	register := 1
	queue := aq.New()

	for _, line := range strings.Split(input, "\n") {
		if slices.Contains(notableCycles, cycle) {
			sum += cycle * register
		}
		cmd := strings.Split(line, " ")
		queue.Enqueue(0)
		if cmd[0] == "addx" {
			val, _ := strconv.Atoi(cmd[1])
			queue.Enqueue(val)
		}
		val, exists := queue.Dequeue()
		if exists {
			register += val.(int)
		}
		cycle++
	}
	for _, val := range queue.Values() {
		if slices.Contains(notableCycles, cycle) {
			sum += cycle * register
		}
		register += val.(int)
		cycle++
	}

	fmt.Println(sum)
}

func partTwo() {
	screen := [240]bool{}
	cycle := 1
	register := 1
	queue := aq.New()

	for _, line := range strings.Split(input, "\n") {
		fmt.Printf("Cycle: %v;Register: %v;Visible: %v\n", cycle, register, math.Abs(float64((cycle-1)%40-register)) <= 1)
		screen[cycle-1] = math.Abs(float64((cycle-1)%40-register)) <= 1
		cmd := strings.Split(line, " ")
		queue.Enqueue(0)
		if cmd[0] == "addx" {
			val, _ := strconv.Atoi(cmd[1])
			queue.Enqueue(val)
		}
		val, exists := queue.Dequeue()
		if exists {
			register += val.(int)
		}
		cycle++
	}
	for _, val := range queue.Values() {
		screen[cycle-1] = math.Abs(float64((cycle-1)%40-register)) <= 1
		register += val.(int)
		cycle++
	}

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			if screen[row*40+col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	//partOne()
	partTwo()
}
