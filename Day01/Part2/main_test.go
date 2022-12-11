
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   getElfsWithMostCalories
/// ---------------------------------------------------------------------------

func Test_getElfsWithMostCalories_emptyElfList(t *testing.T) {
	input := []Elf{}
	expected := []Elf{}

	output := getElfsWithMostCalories(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_getElfsWithMostCalories_TopThreeFromOneElfList(t *testing.T) {
	input := []Elf{
		CreateElfWithFood([]string{"1000"}),
	}
	expected := []Elf{
		CreateElfWithFood([]string{"1000"}),
	}

	output := getElfsWithMostCalories(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_getElfsWithMostCalories_TopOneFromTwoElfList(t *testing.T) {
	input := []Elf{
		CreateElfWithFood([]string{"1000"}),
		CreateElfWithFood([]string{"2000"}),
	}
	expected := []Elf{
		CreateElfWithFood([]string{"2000"}),
	}

	output := getElfsWithMostCalories(input, 1)

	assert.Equal(t, expected, output, "should match")
}

func Test_getElfsWithMostCalories_TopThreeFromSixElfList(t *testing.T) {
	input := []Elf{
		CreateElfWithFood([]string{"1000"}),
		CreateElfWithFood([]string{"2000"}),
		CreateElfWithFood([]string{"3000"}),
		CreateElfWithFood([]string{"4000"}),
		CreateElfWithFood([]string{"5000"}),
		CreateElfWithFood([]string{"6000"}),
	}
	expected := []Elf{
		CreateElfWithFood([]string{"4000"}),
		CreateElfWithFood([]string{"5000"}),
		CreateElfWithFood([]string{"6000"}),
	}

	output := getElfsWithMostCalories(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_getElfsWithMostCalories_TopThreeFromSixElfListInputReverse(t *testing.T) {
	input := []Elf{
		CreateElfWithFood([]string{"6000"}),
		CreateElfWithFood([]string{"5000"}),
		CreateElfWithFood([]string{"4000"}),
		CreateElfWithFood([]string{"3000"}),
		CreateElfWithFood([]string{"2000"}),
		CreateElfWithFood([]string{"1000"}),
	}
	expected := []Elf{
		CreateElfWithFood([]string{"4000"}),
		CreateElfWithFood([]string{"5000"}),
		CreateElfWithFood([]string{"6000"}),
	}

	output := getElfsWithMostCalories(input, 3)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_parseInput_emptyList(t *testing.T) {
	input := []string{}
	expected := []Elf{}

	output := parseInput(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_parseInput_oneElfWithOneItem(t *testing.T) {
	input := []string{"1000"}
	expected := []Elf{
		CreateElfWithFood([]string{"1000"}),
	}

	output := parseInput(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_parseInput_twoElfsWithOneItemEach(t *testing.T) {
	input := []string{"1000", "", "2000"}
	expected := []Elf{
		CreateElfWithFood([]string{"1000"}),
		CreateElfWithFood([]string{"2000"}),
	}

	output := parseInput(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_parseInput_twoElfsWithDifferentCountOfItemsEach(t *testing.T) {
	input := []string{"1000", "", "2000", "3000"}
	expected := []Elf{
		CreateElfWithFood([]string{"1000"}),
		CreateElfWithFood([]string{"2000", "3000"}),
	}

	output := parseInput(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	expected := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}
