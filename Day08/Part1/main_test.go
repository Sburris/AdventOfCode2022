package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   getOverlapWorkordersCount
/// ---------------------------------------------------------------------------

// func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
// 	input := getInput(TEST_INPUT_FILE)
// 	expected := 2

// 	output := getOverlapWorkordersCount(input)

// 	assert.Equal(t, expected, output, "should match")
// }

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
