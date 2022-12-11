package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	Food []int
}

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	elfs := parseInput(input)

	highestCaloies := 0
	for _, elf := range elfs {
		totalCalories := elf.TotalCalories()
		if totalCalories > highestCaloies {
			highestCaloies = totalCalories
		}
	}

	fmt.Print(highestCaloies)
}

func (elf *Elf) TotalCalories() int {
	sum := 0
	for _, item := range elf.Food {
		sum += item
	}

	return sum
}

func parseInput(input []string) []Elf {
	elfs := make([]Elf, 0)

	food := make([]int, 0)
	for _, foodItem := range input {
		// if empty line, create elf with existing food list
		if foodItem == "" {
			if len(food) > 0 {
				newElf := Elf{
					Food: food,
				}
				elfs = append(elfs, newElf)

				// create new empty list for next elf
				food = make([]int, 0)
			}
		} else {
			// Food Item identified
			calories, err := strconv.Atoi(foodItem)
			checkErr(err)

			food = append(food, calories)
		}
	}

	// process last item
	if len(food) > 0 {
		newElf := Elf{
			Food: food,
		}
		elfs = append(elfs, newElf)
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