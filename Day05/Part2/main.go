package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Standard CS Stack
type Stack []string

type Ship struct {
	Cargo []Stack
}

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	fileContents := getInput(inputFile)
	cargo, insturctions := parseInput(fileContents)

	ship := CreateShip(cargo)
	ship.processInstructions(insturctions)
	output := ship.getTopContainers()

	fmt.Print(output)
}

/// ---------------------------------------------------------------------------
///  Ship
/// ---------------------------------------------------------------------------

func CreateShip (cargo []string) Ship {
	newShip := Ship{}

	// Determine number of containers
	lastIndex := len(cargo)-1
	stacks := cargo[lastIndex]
	stacksCount := len(stacks) / 4
	i := 0
	for i <= stacksCount {
		newShip.Cargo = append(newShip.Cargo, CreateStack())
		i += 1
	}

	// Remove stack index
	cargo = cargo[:len(cargo)-1]

	i = len(cargo) - 1
	for i >= 0 {
		line := cargo[i]
		x := 0
		currentContainer := 0
		for x < len(line) {
			item := string(line[x + 1])
			if strings.Trim(item, " ") != "" {
				newShip.Cargo[currentContainer].Push(item)
			}
			currentContainer += 1
			x += 4
		}
		i -= 1
	}

	return newShip
}

func (ship *Ship) getTopContainers() string {
	output := ""

	for _, stack := range ship.Cargo {
		topItem, _ := stack.Pop(1)
		output += topItem
	}

	return output
}

func (ship *Ship) processInstructions(instructions []string) {
	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		count, _ := strconv.Atoi(parts[1])
		fromStack, _ := strconv.Atoi(parts[3])
		toStack, _ := strconv.Atoi(parts[5])

		item, _ := ship.Cargo[fromStack-1].Pop(count)
		ship.Cargo[toStack-1].Push(item)
		count -= 1
	}
}

/// ---------------------------------------------------------------------------
///   Handle Input
/// ---------------------------------------------------------------------------

// Parse out the difference sections of input: Cargo Container and Instructions
func parseInput(input []string) ([]string, []string) {
	// Find empty row

	emptyRow := -1
	for index, line := range input{
		if line == "" {
			emptyRow = index
			break
		}
	}

	cargo := input[:emptyRow]
	instructions := input[emptyRow+1:]

	return cargo, instructions
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

/// ---------------------------------------------------------------------------
///   Stack
/// ---------------------------------------------------------------------------

func CreateStack() Stack {
	return Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(value string) {
	for _, item := range value {
		*s = append(*s, string(item))
	}
}

func (s *Stack) Pop(amount int) (string, bool) {
	if s.IsEmpty() == true {
		return "", false
	}
	index := len(*s) - amount
	element := strings.Join((*s)[index:], "")
	*s = (*s)[:index]
	return element, true
}
