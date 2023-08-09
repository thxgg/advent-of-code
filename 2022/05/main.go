package main

import (
	_ "embed"
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	_ "github.com/emirpasic/gods/stacks/arraystack"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne() {
	rows := strings.Split(input, "\n")

	// Read board
	numOfCols := (len(rows[0]) + 1) / 4
	cols := make([]*arraystack.Stack, numOfCols)
	for col := range cols {
		cols[col] = arraystack.New()
	}

	pegRow := 0

	for i := 0; !strings.HasPrefix(rows[i], " 1 "); i++ {
		pegRow++
	}

	for i := pegRow - 1; i >= 0; i-- {
		row := rows[i]
		for col := 0; col < numOfCols; col++ {
			if row[col*4+1] != ' ' {
				cols[col].Push(string(row[col*4+1]))
			}
		}
	}

	re := regexp.MustCompile("[0-9]+")
	// Read instructions
	for i := pegRow + 2; i < len(rows); i++ {
		row := rows[i]
		move := re.FindAllString(row, -1)
		n, _ := strconv.Atoi(move[0])
		src, _ := strconv.Atoi(move[1])
		tar, _ := strconv.Atoi(move[2])
		for j := 0; j < n; j++ {
			v, _ := cols[src-1].Pop()
			cols[tar-1].Push(v)
		}
	}

	// Output
	for i := 0; i < numOfCols; i++ {
		v, _ := cols[i].Peek()
		fmt.Printf("%v", v)
	}
}

func partTwo() {
	rows := strings.Split(input, "\n")

	// Read board
	numOfCols := (len(rows[0]) + 1) / 4
	cols := make([]*arraystack.Stack, numOfCols)
	for col := range cols {
		cols[col] = arraystack.New()
	}

	pegRow := 0

	for i := 0; !strings.HasPrefix(rows[i], " 1 "); i++ {
		pegRow++
	}

	for i := pegRow - 1; i >= 0; i-- {
		row := rows[i]
		for col := 0; col < numOfCols; col++ {
			if row[col*4+1] != ' ' {
				cols[col].Push(string(row[col*4+1]))
			}
		}
	}

	re := regexp.MustCompile("[0-9]+")
	// Read instructions
	for i := pegRow + 2; i < len(rows); i++ {
		row := rows[i]
		move := re.FindAllString(row, -1)
		n, _ := strconv.Atoi(move[0])
		src, _ := strconv.Atoi(move[1])
		tar, _ := strconv.Atoi(move[2])
		tmp := arraystack.New()
		for j := 0; j < n; j++ {
			v, _ := cols[src-1].Pop()
			tmp.Push(v)
		}
		for j := 0; j < n; j++ {
			v, _ := tmp.Pop()
			cols[tar-1].Push(v)
		}
	}

	// Output
	for i := 0; i < numOfCols; i++ {
		v, _ := cols[i].Peek()
		fmt.Printf("%v", v)
	}
}

func main() {
	//partOne()
	partTwo()
}
