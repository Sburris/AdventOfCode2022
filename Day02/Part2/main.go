package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock int = 1
	Paper    = 2
	Sissors  = 3

	Won      = 4
	Draw     = 5
	Lost     = 6
)

// Coorisponding to the const/enums above
var Points = [7]int {
	-1,  // Unknown/default
	1,   // Shape Rock
	2,   // Shape Paper
	3,   // Shape Sissors
	6,   // Round Won
	3,   // Round Draw
	0,   // Round Lost
}

type Match struct {
	PlayerOneShape int
	PlayerTwoShape int

	PlayerOneScore int
	PlayerTwoScore int
}

var lookupShapeToWin = map[int]int {
	Rock: Paper,
	Paper: Sissors,
	Sissors: Rock,
}

var lookupShapeToLose = map[int]int {
	Rock: Sissors,
	Paper: Rock,
	Sissors: Paper,
}

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	matches := getMatches(inputFile)

	sumPlayerOne := 0
	sumPlayerTwo := 0
	for _, match := range matches {
		sumPlayerOne += match.PlayerOneScore
		sumPlayerTwo += match.PlayerTwoScore
	}

	// 10518
	fmt.Printf("Player One Score: %d\n", sumPlayerOne)

	// 13886
	fmt.Printf("Player Two Score: %d\n", sumPlayerTwo)
}

// Match constructor
func CreateMatch(playerOneShape int, playerTwoDesiredOutcome int) Match {
	newMatch := Match{}

	newMatch.PlayerOneShape = playerOneShape

	shapeToPlay := determineShapeForOutcome(playerOneShape, playerTwoDesiredOutcome)
	newMatch.PlayerTwoShape = shapeToPlay

	newMatch.PlayerOneScore, newMatch.PlayerTwoScore = getPointsForMatch(playerOneShape, shapeToPlay)

	return newMatch
}

// Determine what shape player two should use to get the desired outcome
func determineShapeForOutcome(opponentShape int, desiredOutcome int) int {
	switch desiredOutcome {
	case Draw:
		return opponentShape
	case Won:
		return lookupShapeToWin[opponentShape]
	case Lost:
		return lookupShapeToLose[opponentShape]
	default:
		return -1
	}
}

// Determine who won, lost, or drawed
func getMatchResults(playerOneShape int, playerTwoShape int) (int, int) {
	if playerOneShape == playerTwoShape {
		return Draw, Draw
	}

	if playerOneShape == Rock {
		if playerTwoShape == Paper {
			return Lost, Won
		} else { // Player Two Sissors
			return Won, Lost
		}
	}

	if playerOneShape == Paper {
		if playerTwoShape == Rock {
			return Won, Lost
		} else { // Player Two Sissors
			return Lost, Won
		}
	}

	if playerOneShape == Sissors {
		if playerTwoShape == Rock {
			return Lost, Won
		} else { // Player Two Paper
			return Won, Lost
		}
	}

	// Error
	return -1, -1
}

// Determine the points for the match outcome
func getPointsForMatchResults(result int) int {
	if result < 4 || result > 6 {
		result = 0
	}
	return Points[result]
}

// Determine the points for the shape played
func getPointsForShape(shape int) int {
	if shape <= 0 || shape > 3 {
		shape = 0
	}
	return Points[shape]
}

// Determine the total points each player gets for the match
func getPointsForMatch(playerOneShape int, playerTwoShape int) (int, int) {
	playerOneResults, playerTwoResults := getMatchResults(playerOneShape, playerTwoShape)

	playerOneResultsPoints := getPointsForMatchResults(playerOneResults)
	playerTwoResultsPoints := getPointsForMatchResults(playerTwoResults)

	playerOneShapePoints := getPointsForShape(playerOneShape)
	playerTwoShapePoints := getPointsForShape(playerTwoShape)

	playerOneTotal := playerOneResultsPoints + playerOneShapePoints
	playerTwoTotal := playerTwoResultsPoints + playerTwoShapePoints

	return playerOneTotal, playerTwoTotal
}

var converstionLookup = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Sissors,
	"X": Lost,
	"Y": Draw,
	"Z": Won,
}

func convertInputShapeToProgramShape(input string) int {
	return converstionLookup[input]
}

func getMatches(filename string) []Match {
	input := getInput(filename)
	matches := parseInput(input)
	return matches
}

func parseInput(input []string) []Match {
	matches := make([]Match, 0)

	for _, line := range input {
		splitResults := strings.Split(line, " ")
		playerOneShape := convertInputShapeToProgramShape(splitResults[0])
		playerTwoDesiredOutcome := convertInputShapeToProgramShape(splitResults[1])

		newMatch := CreateMatch(playerOneShape, playerTwoDesiredOutcome)
		matches = append(matches, newMatch)
	}

	return matches
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
