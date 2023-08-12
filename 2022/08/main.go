package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Tree struct {
	height      int
	visible     bool
	scenicScore int
}

func partOne() {
	lines := strings.Split(input, "\n")
	forest := make([][]Tree, len(lines))
	for row, line := range lines {
		treeLine := make([]Tree, len(line))
		for col, tree := range strings.Split(line, "") {
			height, _ := strconv.Atoi(tree)
			treeLine[col] = Tree{height: height, visible: false}
		}
		forest[row] = treeLine
	}

	rows := len(forest)
	cols := len(forest[0])

	// Left parse
	for row := 0; row < rows; row++ {
		maxHeight := math.MinInt
		for col := 0; col < cols; col++ {
			tree := &forest[row][col]
			if maxHeight < tree.height {
				maxHeight = max(maxHeight, tree.height)
				tree.visible = true
			}
		}
	}

	// Right parse
	for row := 0; row < rows; row++ {
		maxHeight := math.MinInt
		for col := cols - 1; col >= 0; col-- {
			tree := &forest[row][col]
			if maxHeight < tree.height {
				maxHeight = max(maxHeight, tree.height)
				tree.visible = true
			}
		}
	}

	// Up parse
	for col := 0; col < cols; col++ {
		maxHeight := math.MinInt
		for row := 0; row < rows; row++ {
			tree := &forest[row][col]
			if maxHeight < tree.height {
				maxHeight = max(maxHeight, tree.height)
				tree.visible = true
			}
		}
	}

	// Down parse
	for col := 0; col < cols; col++ {
		maxHeight := math.MinInt
		for row := rows - 1; row >= 0; row-- {
			tree := &forest[row][col]
			if maxHeight < tree.height {
				maxHeight = max(maxHeight, tree.height)
				tree.visible = true
			}
		}
	}

	// Counting
	visible := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			tree := forest[row][col]
			if tree.visible {
				visible++
			}
		}
	}

	fmt.Println(visible)
}

func partTwo() {
	lines := strings.Split(input, "\n")
	forest := make([][]Tree, len(lines))
	for row, line := range lines {
		treeLine := make([]Tree, len(line))
		for col, tree := range strings.Split(line, "") {
			height, _ := strconv.Atoi(tree)
			treeLine[col] = Tree{height: height, visible: false}
		}
		forest[row] = treeLine
	}

	rows := len(forest)
	cols := len(forest[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// Left parse
			left := 1
			for i := col - 1; i > 0; i-- {
				if forest[row][i].height < forest[row][col].height {
					left++
				} else {
					break
				}
			}
			// Right parse
			right := 1
			for i := col + 1; i < cols-1; i++ {
				if forest[row][i].height < forest[row][col].height {
					right++
				} else {
					break
				}
			}
			// Up parse
			up := 1
			for i := row - 1; i > 0; i-- {
				if forest[i][col].height < forest[row][col].height {
					up++
				} else {
					break
				}
			}
			// Down parse
			down := 1
			for i := row + 1; i < rows-1; i++ {
				if forest[i][col].height < forest[row][col].height {
					down++
				} else {
					break
				}
			}
			// Calculate scenic score
			forest[row][col].scenicScore = left * right * up * down
		}
	}

	// Query
	var maxScenicScore *Tree = nil
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			tree := &forest[row][col]
			if maxScenicScore == nil || maxScenicScore.scenicScore < tree.scenicScore {
				maxScenicScore = tree
			}
		}
	}

	fmt.Println(maxScenicScore.scenicScore)
}

func main() {
	//partOne()
	partTwo()
}
