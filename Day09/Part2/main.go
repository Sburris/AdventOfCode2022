package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rope struct {
	Knots       []RopeKnot
	TailHistory map[string]bool
}

type RopeKnot struct {
	x int
	y int
}

func main() {
	inputFile := "input.txt"
	// inputFile := "test_data/sample_input.txt"
	input := getInput(inputFile)

	rope := CreateRope()

	for _, instruction := range input {
		fmt.Println(instruction)
		rope.MoveInstruction(instruction)
	}

	fmt.Printf("Tail Travel History: %d\n", len(rope.TailHistory))
}

func CreateRope() Rope {
	newRope := Rope{}
	newRope.Knots = []RopeKnot{
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
		{x: 0, y: 0},
	}
	newRope.TailHistory = make(map[string]bool)

	return newRope
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
		fmt.Printf("%s %d\n", direction, currentCount)
		rope.move(direction)
		currentCount += 1
	}
}

func (rope *Rope) move(direction string) {
	for index := range rope.Knots {

		// Head
		if index == 0 {
			rope.Knots[index].x += MoveDirection[direction][0]
			rope.Knots[index].y += MoveDirection[direction][1]
			fmt.Printf("%d (%d, %d)\n", index, rope.Knots[index].x, rope.Knots[index].y)
			continue
		}

		// Trailing Knots
		xDiff := Abs(rope.Knots[index].x - rope.Knots[index-1].x)
		yDiff := Abs(rope.Knots[index].y - rope.Knots[index-1].y)
		if xDiff > 1 || yDiff > 1 {
			if xDiff == 2 {
				rope.Knots[index].x = Average(rope.Knots[index].x, rope.Knots[index-1].x)
			} else if xDiff == 1 {
				rope.Knots[index].x = rope.Knots[index-1].x
			}

			if yDiff == 2 {
				rope.Knots[index].y = Average(rope.Knots[index].y, rope.Knots[index-1].y)
			} else if yDiff == 1 {
				rope.Knots[index].y = rope.Knots[index-1].y
			}
		}

		fmt.Printf("%d (%d, %d)\n", index, rope.Knots[index].x, rope.Knots[index].y)

		// Record Tail position
		if index == 9 {
			location := fmt.Sprintf("%d,%d", rope.Knots[index].x, rope.Knots[index].y)
			rope.TailHistory[location] = true
		}

	}
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func Average(first int, second int) int {
	return (first + second) / 2
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
