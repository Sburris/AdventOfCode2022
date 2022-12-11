package main

import (
	"bufio"
	"fmt"
	"os"
)

var changeDir = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Forest struct {
	Trees      []string
	sideLength int
}

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)

	forest := CreateForest(input)
	highestScenicScore := forest.GetHighestScenicScore()

	fmt.Println(highestScenicScore)
}

func CreateForest(input []string) Forest {
	newForest := Forest{}

	newForest.Trees = input
	newForest.sideLength = len(input)

	return newForest
}

func (forest *Forest) GetHighestScenicScore() int {
	xIndex := 0
	highestScenicScore := 0
	for xIndex < forest.sideLength {
		yIndex := 0
		for yIndex < forest.sideLength {
			directionIndex := 0
			scenicScore := 1
			for directionIndex < 4 {
				score := forest.getScenicScore(xIndex, yIndex, directionIndex)
				scenicScore *= score
				directionIndex += 1
			}

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}

			yIndex += 1
		}
		xIndex += 1
	}
	return highestScenicScore
}

func (forest *Forest) getScenicScore(x int, y int, dir int) int {
	scenicScore := 0
	currentTreeSize := forest.getTree(x, y)
	for true {
		x += changeDir[dir][0]
		y += changeDir[dir][1]
		if forest.isOutOfBounds(x, y) {
			break
		}
		scenicScore += 1
		treeSize := forest.getTree(x, y)
		if treeSize >= currentTreeSize {
			break
		}
	}
	return scenicScore
}

func (forest *Forest) isOutOfBounds(x int, y int) bool {
	return x < 0 || x >= forest.sideLength || y < 0 || y >= forest.sideLength
}

func (forest *Forest) getTree(x int, y int) int {
	temp := int(forest.Trees[y][x])
	return temp - 48
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
