package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	// inputFile := "input.txt"
	inputFile := "test_data/sample_input.txt"

	// input := getInput(inputFile)
	// overlaps := getOverlapWorkordersCount(input)

	fmt.Print(inputFile)
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
