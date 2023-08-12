package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Knot struct {
	x    int
	y    int
	next *Knot
}

func closeDistance(head *Knot, tail *Knot) {
	distX := head.x - tail.x
	distY := head.y - tail.y
	if distX == 2 {
		tail.x++
		tail.y += distY
	} else if distX == -2 {
		tail.x--
		tail.y += distY
	} else if distY == 2 {
		tail.x += distX
		tail.y++
	} else if distY == -2 {
		tail.x += distX
		tail.y--
	}
}

func move(head *Knot, visited map[int]map[int]bool, cmd string) {
	x := strings.Split(cmd, " ")
	dir := x[0]
	n, _ := strconv.Atoi(x[1])
	for i := 0; i < n; i++ {
		switch dir {
		case "L":
			head.x--
		case "U":
			head.y++
		case "R":
			head.x++
		case "D":
			head.y--
		}
		knot := head
		for knot.next != nil {
			closeDistance(knot, knot.next)
			knot = knot.next
		}
		if visited[knot.x] == nil {
			visited[knot.x] = make(map[int]bool)
		}
		visited[knot.x][knot.y] = true
	}
}

func partOne() {
	tail := Knot{}
	head := Knot{next: &tail}
	visited := make(map[int]map[int]bool)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		move(&head, visited, line)
	}

	count := 0
	for _, row := range visited {
		count += len(row)
	}
	fmt.Println(count)
}

func partTwo() {
	knots := make([]*Knot, 9)
	for i := 0; i < 9; i++ {
		knots[i] = &Knot{}
		if i != 0 {
			knots[i-1].next = knots[i]
		}
	}
	visited := make(map[int]map[int]bool)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		move(knots[0], visited, line)
	}

	count := 0
	for _, row := range visited {
		count += len(row)
	}
	fmt.Println(count)
}

func main() {
	//partOne()
	partTwo()
}
