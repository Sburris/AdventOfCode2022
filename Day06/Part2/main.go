package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	marker, index := getMarker(input, 4)

	fmt.Println(marker)
	fmt.Println(index + len(marker))
}

func getMarker(input string, markerLength int) (string, int) {
	lookup := make(map[string]int)

	currentLength := 0
	pointer := 0
	inputLength := len(input)
	for pointer < inputLength {
		letter := string(input[pointer])
		value, found := lookup[letter]
		if found == true && (pointer > 0 && value != -1) {
			temp := pointer - currentLength
			for temp <= value {
				lookup[string(input[temp])] = -1
				currentLength -= 1
				temp += 1
			}
			lookup[letter] = pointer
			currentLength += 1
		} else {
			lookup[letter] = pointer
			currentLength += 1
		}
		pointer += 1

		if currentLength == markerLength {
			start := pointer - markerLength
			end := pointer
			marker := input[start:end]
			return marker, start
		}
	}

	return "", -1
}

// Read in raw data so we can process it
func getInput(filename string) string {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return line
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
