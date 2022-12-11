package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile := "input.txt"
	// inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	// overlaps := getOverlapWorkordersCount(input)

	sideLength := len(input)

	treeVisibility := make([][]bool, sideLength, sideLength)
	row := 0
	for row < sideLength {
		treeVisibility[row] = make([]bool, sideLength, sideLength)
		row += 1
	}

	visibiltyCount := 0

	rowColumnIndex := 0
	for rowColumnIndex < sideLength {
		rowLeftHeight := -1
		rowRightHeight := -1
		columnLeftHeight := -1
		columnRightHeight := -1
		lookIndex := 0
		for lookIndex < sideLength {
			rowLeftTree := int(input[rowColumnIndex][lookIndex] - 48)
			if rowLeftTree > rowLeftHeight {
				rowLeftHeight = rowLeftTree
				if treeVisibility[rowColumnIndex][lookIndex] == false {
					visibiltyCount += 1
					treeVisibility[rowColumnIndex][lookIndex] = true
				}
			}

			rowRightTree := int(input[rowColumnIndex][sideLength-lookIndex-1] - 48)
			if rowRightTree > rowRightHeight {
				rowRightHeight = rowRightTree
				if treeVisibility[rowColumnIndex][sideLength-lookIndex-1] == false {
					visibiltyCount += 1
					treeVisibility[rowColumnIndex][sideLength-lookIndex-1] = true
				}
			}

			columnLeftTree := int(input[lookIndex][rowColumnIndex] - 48)
			if columnLeftTree > columnLeftHeight {
				columnLeftHeight = columnLeftTree
				if treeVisibility[lookIndex][rowColumnIndex] == false {
					visibiltyCount += 1
					treeVisibility[lookIndex][rowColumnIndex] = true
				}
			}

			columnRightTree := int(input[sideLength-lookIndex-1][rowColumnIndex] - 48)
			if columnRightTree > columnRightHeight {
				columnRightHeight = columnRightTree
				if treeVisibility[sideLength-lookIndex-1][rowColumnIndex] == false {
					visibiltyCount += 1
					treeVisibility[sideLength-lookIndex-1][rowColumnIndex] = true
				}
			}

			lookIndex += 1
		}
		rowColumnIndex += 1
	}

	fmt.Println(visibiltyCount)
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
