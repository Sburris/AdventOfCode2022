package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type RucksackAuditor struct {
	rucksacks []Rucksack
}

type Rucksack struct {
	compartmentOne string
	compartmentTwo string
}

func CreateRucksackAuditor(input []string) RucksackAuditor {
	newRucksackAuditor := RucksackAuditor{}

	for _, line := range input {
		newRucksack := CreateRucksack(line)
		newRucksackAuditor.rucksacks = append(newRucksackAuditor.rucksacks, newRucksack)
	}

	return newRucksackAuditor
}

func (auditor *RucksackAuditor) GetScore() int {
	sum := 0
	for _, ruck := range auditor.rucksacks {
		item := ruck.GetOverlappingItem()
		sum += GetSupplyItemPriority(item)
	}
	return sum
}

func CreateRucksack(input string) Rucksack {
	newRucksack := Rucksack{}

	halfway := len(input) / 2
	newRucksack.compartmentOne = input[:halfway]
	newRucksack.compartmentTwo = input[halfway:]

	return newRucksack
}

func (ruck *Rucksack) GetOverlappingItem() string {
	// Get distinct items from first compartment
	uniqueItemsCompartmentOne := ruck.getDistinctItems(ruck.compartmentOne)
	pattern := fmt.Sprintf("[%s]", uniqueItemsCompartmentOne)
	regex := regexp.MustCompile(pattern)
	match := regex.FindString(ruck.compartmentTwo)

	return match
}

func GetSupplyItemPriority(item string) int {
	value := int(item[0])

	if value > 90 {
		return value - 96
	}
	return value - 38
}

func (ruck *Rucksack) getDistinctItems(items string) string {
	lookup := make(map[rune]int, 0)

	for _, item := range items {
		_, found := lookup[item]
		if found == false {
			lookup[item] = 1
		}
	}

	uniqueItems := ""
	for key := range lookup {
		uniqueItems += string(key)
	}

	return uniqueItems
}

// TODO: Determine Supply Item Priority

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	auditor := CreateRucksackAuditor(input)

	score := auditor.GetScore()

	fmt.Print(score)
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
