package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   Ship.processInstructions
/// ---------------------------------------------------------------------------

func Test_ShipProcessInstructions_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := "CMZ"

	fileContents := getInput(input)
	cargo, insturctions := parseInput(fileContents)

	ship := CreateShip(cargo)
	ship.processInstructions(insturctions)
	output := ship.getTopContainers()

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   Ship.getTopContainers
/// ---------------------------------------------------------------------------

func Test_ShipGetTopContainers_sampleInput(t *testing.T) {
	input := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
	}
	expected := "NDP"

	ship := CreateShip(input)
	output := ship.getTopContainers()

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   CreateShip
/// ---------------------------------------------------------------------------

func Test_CreateShip_sampleInput(t *testing.T) {
	input := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
	}
	expected := Ship{}
	expected.Cargo = append(expected.Cargo, CreateStack())
	expected.Cargo = append(expected.Cargo, CreateStack())
	expected.Cargo = append(expected.Cargo, CreateStack())
	expected.Cargo[0].Push("Z")
	expected.Cargo[0].Push("N")
	expected.Cargo[1].Push("M")
	expected.Cargo[1].Push("C")
	expected.Cargo[1].Push("D")
	expected.Cargo[2].Push("P")

	output := CreateShip(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   parseInput
/// ---------------------------------------------------------------------------

func Test_parseInput_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expectedCargo := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
	}
	expectedInstructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	fileContents := getInput(input)
	outputCargo, outputInsturctions := parseInput(fileContents)

	assert.Equal(t, expectedCargo, outputCargo, "should match")
	assert.Equal(t, expectedInstructions, outputInsturctions, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{
		"    [D]",
		"[N] [C]",
		"[Z] [M] [P]",
		" 1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
