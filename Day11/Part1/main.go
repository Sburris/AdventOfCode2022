package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []int
	Operations      []int
	IsDoubled       bool
	DivisibleBy     int
	Test            []int
	InspectionCount int
}

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	monkeies := ParseInput(input)

	roundCount := 20
	round := 0
	for round < roundCount {
		monkeyCount := len(monkeies)
		monkeyIndex := 0
		for monkeyIndex < monkeyCount {
			for {
				item, nextMonkey := monkeies[monkeyIndex].InspectItem()
				if item == -1 {
					break
				}
				monkeies[nextMonkey].Items = enqueue(monkeies[nextMonkey].Items, item)
			}
			monkeyIndex += 1
		}

		round += 1
	}

	inspectionCounts := make([]int, 0)
	for _, monkey := range monkeies {
		inspectionCounts = append(inspectionCounts, monkey.InspectionCount)
	}

	sort.Ints(inspectionCounts)
	count := len(inspectionCounts)
	num := inspectionCounts[count-2] * inspectionCounts[count-1]

	fmt.Print(num)
}

func (monkey *Monkey) InspectItem() (int, int) {
	if len(monkey.Items) == 0 {
		return -1, -1
	}

	monkey.InspectionCount += 1
	item, items := dequeue(monkey.Items)
	monkey.Items = items

	worryLevel := item
	if monkey.IsDoubled {
		worryLevel *= worryLevel
	} else {
		worryLevel *= monkey.Operations[0]
	}
	worryLevel += monkey.Operations[1]
	worryLevel /= 3
	isDivisible := (worryLevel % monkey.DivisibleBy) == 0
	if isDivisible == true {
		return worryLevel, monkey.Test[0]
	}
	return worryLevel, monkey.Test[1]

}

func CreateMonkey(input []string) Monkey {
	monkey := Monkey{
		InspectionCount: 0,
	}

	// Items
	itemParts := strings.Split(input[0][18:], ", ")
	monkey.Items = make([]int, 0)
	for _, part := range itemParts {
		value, _ := strconv.Atoi(part)
		monkey.Items = enqueue(monkey.Items, value)
	}

	// Operations
	opsParts := strings.Split(input[1][23:], " ")
	opsValue, _ := strconv.Atoi(opsParts[1])
	monkey.Operations = []int{1, 0}
	switch opsParts[0] {
	case "*":
		monkey.IsDoubled = opsParts[1] == "old"

		monkey.Operations[0] = opsValue
	case "+":
		monkey.Operations[1] = opsValue
	default:
		panic("UNKNOWN operation")
	}

	// Divisible By
	divValue, _ := strconv.Atoi(input[2][21:])
	monkey.DivisibleBy = divValue

	// Test
	trueValue, _ := strconv.Atoi(input[3][29:])
	falseValue, _ := strconv.Atoi(input[4][30:])
	monkey.Test = []int{trueValue, falseValue}

	return monkey
}

func ParseInput(input []string) []Monkey {
	monkeies := make([]Monkey, 0)

	monkeyCount := (len(input) / 7) + 1
	index := 0
	for index < monkeyCount {
		start := index * 7
		newMonkey := CreateMonkey(input[start+1 : start+6])
		monkeies = append(monkeies, newMonkey)
		index += 1
	}

	return monkeies
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

func enqueue(queue []int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []int) (int, []int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []int{}
		return element, tmp
	}
	return element, queue[1:]
}
