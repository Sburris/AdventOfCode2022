package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []big.Int
	Operations      []big.Int
	IsDoubled       bool
	DivisibleBy     big.Int
	Test            []int
	InspectionCount int
}

var NegOne *big.Int = big.NewInt(-1)
var Zero *big.Int = big.NewInt(0)

var LCM *big.Int = big.NewInt(9699690) // Least Common Multiple (input file)
//var LCM *big.Int = big.NewInt(96577) // Least Common Multiple (sample file)

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	monkeies := ParseInput(input)

	roundCount := 10000
	round := 0
	for round < roundCount {
		monkeyCount := len(monkeies)
		monkeyIndex := 0
		for monkeyIndex < monkeyCount {
			for {
				item, nextMonkey := monkeies[monkeyIndex].InspectItem()
				if item.Cmp(NegOne) == 0 {
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
	//	num := inspectionCounts[count-2] * inspectionCounts[count-1]
	num := big.NewInt(0)
	num1 := int64(inspectionCounts[count-2])
	num2 := int64(inspectionCounts[count-1])
	num.Mul(big.NewInt(num1), big.NewInt(num2))
	fmt.Print(num)
}

func (monkey *Monkey) InspectItem() (big.Int, int) {
	if len(monkey.Items) == 0 {
		return *NegOne, -1
	}

	monkey.InspectionCount += 1
	item, items := dequeue(monkey.Items)
	monkey.Items = items

	worryLevel := item
	if monkey.IsDoubled {
		worryLevel.Mul(&worryLevel, &worryLevel)
	} else {
		worryLevel.Mul(&worryLevel, &monkey.Operations[0])
	}
	worryLevel.Add(&worryLevel, &monkey.Operations[1])
	worryLevel.Mod(&worryLevel, LCM)
	//worryLevel /= 3
	isDivisible := big.NewInt(0)
	isDivisible.Mod(&worryLevel, &monkey.DivisibleBy)
	if isDivisible.Cmp(big.NewInt(0)) == 0 {
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
	monkey.Items = make([]big.Int, 0)
	for _, part := range itemParts {
		value, _ := strconv.ParseInt(part, 10, 64)
		monkey.Items = enqueue(monkey.Items, *big.NewInt(value))
	}

	// Operations
	opsParts := strings.Split(input[1][23:], " ")
	opsValueInt64, _ := strconv.ParseInt(opsParts[1], 10, 64)
	opsValue := big.NewInt(opsValueInt64)
	monkey.Operations = []big.Int{*big.NewInt(1), *big.NewInt(0)}
	switch opsParts[0] {
	case "*":
		monkey.IsDoubled = opsParts[1] == "old"

		monkey.Operations[0] = *opsValue
	case "+":
		monkey.Operations[1] = *opsValue
	default:
		panic("UNKNOWN operation")
	}

	// Divisible By
	divValue, _ := strconv.ParseInt(input[2][21:], 10, 64)
	monkey.DivisibleBy = *big.NewInt(divValue)

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

func enqueue(queue []big.Int, element big.Int) []big.Int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []big.Int) (big.Int, []big.Int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []big.Int{}
		return element, tmp
	}
	return element, queue[1:]
}
