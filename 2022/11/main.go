package main

import (
	_ "embed"
	"fmt"
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Monkey struct {
	items       *aq.Queue
	operation   func(int) int
	test        func(int) int
	mod         int
	inspections int
}

func operation(op string, a int, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func parseMonkey(monkeyInfo string) Monkey {
	monkey := Monkey{}

	info := strings.Split(monkeyInfo, "\n")

	// Items
	monkey.items = aq.New()
	items := strings.Split(strings.Split(info[1], ": ")[1], ", ")
	for _, item := range items {
		parsedItem, _ := strconv.Atoi(item)
		monkey.items.Enqueue(parsedItem)
	}

	//Operation
	opInfo := strings.Split(strings.Split(info[2], "= ")[1], " ")
	monkey.operation = func(n int) int {
		b := opInfo[2]
		if b == "old" {
			return operation(opInfo[1], n, n)
		} else {
			parsedB, _ := strconv.Atoi(b)
			return operation(opInfo[1], n, parsedB)
		}
	}

	//Test
	divBy, _ := strconv.Atoi(strings.Split(info[3], "by ")[1])
	tarTrue, _ := strconv.Atoi(strings.Split(info[4], "monkey ")[1])
	tarFalse, _ := strconv.Atoi(strings.Split(info[5], "monkey ")[1])
	monkey.mod = divBy
	monkey.test = func(n int) int {
		if n%divBy == 0 {
			return tarTrue
		} else {
			return tarFalse
		}
	}

	return monkey
}

func playRound(monkeys []*Monkey, divBy3 bool, mod int) {
	for _, monkey := range monkeys {
		for !monkey.items.Empty() {
			monkey.inspections++
			item, _ := monkey.items.Dequeue()
			val := item.(int)
			val = monkey.operation(val) % mod
			if divBy3 {
				val /= 3
			}
			monkeys[monkey.test(val)].items.Enqueue(val)
		}
	}
}

func partOne() {
	var monkeys []*Monkey
	mod := 1
	for _, monkeyInfo := range strings.Split(input, "\n\n") {
		parsedMonkey := parseMonkey(monkeyInfo)
		monkeys = append(monkeys, &parsedMonkey)
		mod *= parsedMonkey.mod
	}

	for i := 0; i < 20; i++ {
		playRound(monkeys, true, mod)
	}

	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0] * inspections[1])
}

func partTwo() {
	var monkeys []*Monkey
	mod := 1
	for _, monkeyInfo := range strings.Split(input, "\n\n") {
		parsedMonkey := parseMonkey(monkeyInfo)
		monkeys = append(monkeys, &parsedMonkey)
		mod *= parsedMonkey.mod
	}

	for i := 0; i < 10000; i++ {
		playRound(monkeys, false, mod)
	}

	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Println(inspections[0] * inspections[1])
}

func main() {
	//partOne()
	partTwo()
}
