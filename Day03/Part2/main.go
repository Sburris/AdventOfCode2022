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
	allTiems string
	compartmentOne string
	compartmentTwo string
	itemCount map[string]int
}

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	auditor := CreateRucksackAuditor(input)

	badges := auditor.GetAllBadages()

	sum := 0
	for _, badge := range badges {
		priority := GetSupplyItemPriority(string(badge))
		sum += priority
	}

	fmt.Print(sum)
}

func GetSupplyItemPriority(item string) int {
	value := int(item[0])

	// Determine if this is a lowercase vs uppercase
	if value > 90 {
		return value - 96
	}
	return value - 38
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

func (auditor *RucksackAuditor) GetAllBadages() []string {
	badges := make([]string, 0)

	currentGroup := 0
	for currentGroup < len(auditor.rucksacks) {
		for item := range auditor.rucksacks[currentGroup].itemCount {
			_, found2 := auditor.rucksacks[currentGroup+1].itemCount[item]
			_, found3 := auditor.rucksacks[currentGroup+2].itemCount[item]
			if found2 == true && found3 == true {
				badges = append(badges, item)
			}
		}

		currentGroup += 3
	}

	return badges
}

func CreateRucksack(input string) Rucksack {
	newRucksack := Rucksack{}

	newRucksack.allTiems = input
	halfway := len(input) / 2
	newRucksack.compartmentOne = input[:halfway]
	newRucksack.compartmentTwo = input[halfway:]

	newRucksack.itemCount = make(map[string]int)
	newRucksack.CountInventory()

	return newRucksack
}

func (ruck *Rucksack) CountInventory() {
	for _, i := range ruck.allTiems {
		item := string(i)
		count, found := ruck.itemCount[item]
		if found == true {
			ruck.itemCount[item] = count + 1
		} else {
			ruck.itemCount[item] = 1
		}
	}
}

// Get all unquie items in bag
func (ruck *Rucksack) GetPossibleBadges() map[string]int {
	possibleBadges := make(map[string]int, 0)

	for item, count := range ruck.itemCount {
		if count == 1 {
			possibleBadges[string(item)] = 1
		}
	}

	return possibleBadges
}

func (ruck *Rucksack) GetOverlappingItem() string {
	// Get distinct items from first compartment
	uniqueItemsCompartmentOne := ruck.getDistinctItems(ruck.compartmentOne)
	pattern := fmt.Sprintf("[%s]", uniqueItemsCompartmentOne)
	regex := regexp.MustCompile(pattern)
	match := regex.FindString(ruck.compartmentTwo)

	return match
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
