package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	elfs := getElfData(inputFile)
	topThreeElfs := getElfsWithMostCalories(elfs, 3)
	totalCaloriesCarried := getTotalCaloriesCarried(topThreeElfs)

	// 202585
	fmt.Print(totalCaloriesCarried)
}

func getTotalCaloriesCarried(elfs []Elf) int {
	sum := 0
	for _, elf := range elfs {
		sum += elf.CaloriesCarried
	}
	return sum
}

func getElfsWithMostCalories(elfs []Elf, count int) []Elf {
	// Create copy of list so we are not modifiying the list
	// outside of this function as a side effect
	elfList := []Elf{}
	elfList = append(elfList, elfs...)

	if len(elfList) <= count {
		return elfList
	}

	//Sort Elfs based on Total Calories
	sort.Slice(elfList, func(a, b int) bool {
		return elfList[a].CaloriesCarried < elfList[b].CaloriesCarried
	})

	return elfList[len(elfList)-count:]
}

func getElfData(inputFile string) []Elf {
	input := getInput(inputFile)
	elfs := parseInput(input)
	return elfs
}

func parseInput(input []string) []Elf {
	elfs := make([]Elf, 0)

	currentElf := CreateElf()
	for _, foodItem := range input {
		// if empty line, create elf with existing food list
		if foodItem == "" {
			if currentElf.CaloriesCarried > 0 {
				elfs = append(elfs, currentElf)
				currentElf = CreateElf()
			}
		} else {
			// else add food item to current elf
			currentElf.AddFood(foodItem)
		}
	}

	// process last elf
	if currentElf.CaloriesCarried > 0 {
		elfs = append(elfs, currentElf)
	}

	return elfs
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
