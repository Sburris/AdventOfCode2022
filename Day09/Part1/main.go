package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rope struct {
	Head         RopeEnd
	HeadPrevious RopeEnd
	Tail         RopeEnd
	TailHistory  map[string]bool
}

type RopeEnd struct {
	x int
	y int
}

func main() {
	inputFile := "input.txt"
	// inputFile := "test_data/sample_input.txt"
	input := getInput(inputFile)

	rope := CreateRope()

	for _, instruction := range input {
		rope.MoveInstruction(instruction)
	}

	fmt.Print(len(rope.TailHistory))
}

func CreateRope() Rope {
	return Rope{
		Head: RopeEnd{
			x: 0,
			y: 0,
		},
		Tail: RopeEnd{
			x: 0,
			y: 0,
		},
		TailHistory: make(map[string]bool),
	}
}

var MoveDirection = map[string][]int{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func (rope *Rope) MoveInstruction(instruction string) {
	parts := strings.Split(instruction, " ")
	direction := parts[0]
	times, _ := strconv.Atoi(parts[1])

	currentCount := 0
	for currentCount < times {
		rope.move(direction)
		currentCount += 1
	}
}

func (rope *Rope) move(direction string) {
	rope.moveHead(direction)
	rope.moveTail()
}

func (rope *Rope) moveHead(direction string) {
	rope.HeadPrevious = rope.Head
	rope.Head.x += MoveDirection[direction][0]
	rope.Head.y += MoveDirection[direction][1]
}

func (rope *Rope) moveTail() {
	xDiff := Abs(rope.Head.x - rope.Tail.x)
	yDiff := Abs(rope.Head.y - rope.Tail.y)
	if xDiff > 1 || yDiff > 1 {
		rope.Tail = rope.HeadPrevious
	}
	location := fmt.Sprintf("%d,%d", rope.Tail.x, rope.Tail.y)
	rope.TailHistory[location] = true
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// Read in raw data so we can process it
func getInput(filename string) []string {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
