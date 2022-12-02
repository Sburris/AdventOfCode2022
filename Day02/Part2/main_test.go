
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   convertInputShapeToProgramShape
/// ---------------------------------------------------------------------------

func Test_convertInputShapeToProgramShape_A(t *testing.T) {
	input := "A"
	expected := Rock

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_convertInputShapeToProgramShape_B(t *testing.T) {
	input := "B"
	expected := Paper

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_convertInputShapeToProgramShape_C(t *testing.T) {
	input := "C"
	expected := Sissors

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_convertInputShapeToProgramShape_X(t *testing.T) {
	input := "X"
	expected := Lost

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_convertInputShapeToProgramShape_Y(t *testing.T) {
	input := "Y"
	expected := Draw

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_convertInputShapeToProgramShape_Z(t *testing.T) {
	input := "Z"
	expected := Won

	output := convertInputShapeToProgramShape(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   determineShapeForOutcome
/// ---------------------------------------------------------------------------

func Test_determineShapeForOutcome_RockWon(t *testing.T) {
	inputPlayerOne := Rock
	inputPlayerTwo := Won
	expectedShape := Paper

	shape := determineShapeForOutcome(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedShape, shape, "should match")
}

func Test_determineShapeForOutcome_PaperLost(t *testing.T) {
	inputPlayerOne := Paper
	inputPlayerTwo := Lost
	expectedShape := Rock

	shape := determineShapeForOutcome(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedShape, shape, "should match")
}

func Test_determineShapeForOutcome_SissorDraw(t *testing.T) {
	inputPlayerOne := Sissors
	inputPlayerTwo := Draw
	expectedShape := Sissors

	shape := determineShapeForOutcome(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedShape, shape, "should match")
}

/// ---------------------------------------------------------------------------
///   getMatchResults
/// ---------------------------------------------------------------------------

func Test_getMatchResults_RockRock(t *testing.T) {
	inputPlayerOne := Rock
	inputPlayerTwo := Rock
	expectedPlayerOne := Draw
	expectedPlayerTwo := Draw

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_RockPaper(t *testing.T) {
	inputPlayerOne := Rock
	inputPlayerTwo := Paper
	expectedPlayerOne := Lost
	expectedPlayerTwo := Won

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_RockSissors(t *testing.T) {
	inputPlayerOne := Rock
	inputPlayerTwo := Sissors
	expectedPlayerOne := Won
	expectedPlayerTwo := Lost

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_PaperRock(t *testing.T) {
	inputPlayerOne := Paper
	inputPlayerTwo := Rock
	expectedPlayerOne := Won
	expectedPlayerTwo := Lost

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_PaperPaper(t *testing.T) {
	inputPlayerOne := Paper
	inputPlayerTwo := Paper
	expectedPlayerOne := Draw
	expectedPlayerTwo := Draw

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_PaperSissors(t *testing.T) {
	inputPlayerOne := Paper
	inputPlayerTwo := Sissors
	expectedPlayerOne := Lost
	expectedPlayerTwo := Won

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_SissorsRock(t *testing.T) {
	inputPlayerOne := Sissors
	inputPlayerTwo := Rock
	expectedPlayerOne := Lost
	expectedPlayerTwo := Won

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_SissorsPaper(t *testing.T) {
	inputPlayerOne := Sissors
	inputPlayerTwo := Paper
	expectedPlayerOne := Won
	expectedPlayerTwo := Lost

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getMatchResults_SissorsSissors(t *testing.T) {
	inputPlayerOne := Sissors
	inputPlayerTwo := Sissors
	expectedPlayerOne := Draw
	expectedPlayerTwo := Draw

	outputPlayerOne, outputPlayerTwo := getMatchResults(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

/// ---------------------------------------------------------------------------
///   getPointsForMatchResults
/// ---------------------------------------------------------------------------

func Test_getPointsForMatchResults_OutOfBoundsLow(t *testing.T) {
	input := -100
	expected := -1

	output := getPointsForMatchResults(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForMatchResults_OutOfBoundsHigh(t *testing.T) {
	input := 999999
	expected := -1

	output := getPointsForMatchResults(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForMatchResults_Win(t *testing.T) {
	input := Won
	expected := 6

	output := getPointsForMatchResults(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForMatchResults_Lost(t *testing.T) {
	input := Lost
	expected := 0

	output := getPointsForMatchResults(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForMatchResults_Draw(t *testing.T) {
	input := Won
	expected := 6

	output := getPointsForMatchResults(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getPointsForHand
/// ---------------------------------------------------------------------------

func Test_getPointsForShape_OutOfRangeLow(t *testing.T) {
	input := -100
	expected := -1

	output := getPointsForShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForShape_OutOfRangeHigh(t *testing.T) {
	input := 99999
	expected := -1

	output := getPointsForShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForShape_Rock(t *testing.T) {
	input := Rock
	expected := 1

	output := getPointsForShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForShape_Paper(t *testing.T) {
	input := Paper
	expected := 2

	output := getPointsForShape(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_getPointsForShape_Sissors(t *testing.T) {
	input := Sissors
	expected := 3

	output := getPointsForShape(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getPointsForMatch
/// ---------------------------------------------------------------------------

func Test_getPointsForMatch_WonLost(t *testing.T) {
	inputPlayerOne := Sissors
	inputPlayerTwo := Paper
	expectedPlayerOne := 9
	expectedPlayerTwo := 2

	outputPlayerOne, outputPlayerTwo := getPointsForMatch(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getPointsForMatch_LostWon(t *testing.T) {
	inputPlayerOne := Paper
	inputPlayerTwo := Sissors
	expectedPlayerOne := 2
	expectedPlayerTwo := 9

	outputPlayerOne, outputPlayerTwo := getPointsForMatch(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}

func Test_getPointsForMatch_DrawDraw(t *testing.T) {
	inputPlayerOne := Rock
	inputPlayerTwo := Rock
	expectedPlayerOne := 4
	expectedPlayerTwo := 4

	outputPlayerOne, outputPlayerTwo := getPointsForMatch(inputPlayerOne, inputPlayerTwo)

	assert.Equal(t, expectedPlayerOne, outputPlayerOne, "should match")
	assert.Equal(t, expectedPlayerTwo, outputPlayerTwo, "should match")
}
