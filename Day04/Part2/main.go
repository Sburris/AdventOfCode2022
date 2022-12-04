package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	overlaps := getOverlapWorkordersCount(input)

	fmt.Print(overlaps)
}

func getOverlapWorkordersCount(input []string) int {
	sum := 0

	for _, line := range input {
		e1Start, e1End, e2Start, e2End := getWorkOrders(line)
		overlap := doesContainOverlap(e1Start, e1End, e2Start, e2End)
		if overlap == true {
			sum += 1
		}
	}

	return sum
}

func getWorkOrders(line string) (int, int, int , int) {
	workOrders := strings.Split(line, ",")
	firstStart, firstEnd := convertWorkOrder(workOrders[0])
	secondStart, secondEnd := convertWorkOrder(workOrders[1])

	return firstStart, firstEnd, secondStart, secondEnd
}

func convertWorkOrder(workorder string) (int, int) {
	parts := strings.Split(workorder, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}

func doesContainSubset(e1Start int, e1End int, e2Start int, e2End int) bool {
	if (e1Start <= e2Start && e1End >= e2End) || (e2Start <= e1Start && e2End >= e1End) {
		return true
	}
	return false
}

func doesContainOverlap(e1Start int, e1End int, e2Start int, e2End int) bool {
	if (e1Start <= e2End && e1End >= e2Start) || (e2Start <= e1End && e2End >= e1Start) {
		return true
	}
	return false
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
