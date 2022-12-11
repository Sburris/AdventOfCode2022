package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   getMarker
/// ---------------------------------------------------------------------------

func Test_getMarker_sample1(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expectedMarker := "jpqm"
	expectedIndex := 3

	outputMarker, ouputIndex := getMarker(input, 4)

	assert.Equal(t, expectedMarker, outputMarker, "should match")
	assert.Equal(t, expectedIndex, ouputIndex, "should match")
}

func Test_getMarker_sample2(t *testing.T) {
	input := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	expectedMarker := "vwbj"
	expectedIndex := 1

	outputMarker, ouputIndex := getMarker(input, 4)

	assert.Equal(t, expectedMarker, outputMarker, "should match")
	assert.Equal(t, expectedIndex, ouputIndex, "should match")
}

func Test_getMarker_sample3(t *testing.T) {
	input := "nppdvjthqldpwncqszvftbrmjlhg"
	expectedMarker := "pdvj"
	expectedIndex := 2

	outputMarker, ouputIndex := getMarker(input, 4)

	assert.Equal(t, expectedMarker, outputMarker, "should match")
	assert.Equal(t, expectedIndex, ouputIndex, "should match")
}

func Test_getMarker_sample4(t *testing.T) {
	input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	expectedMarker := "rfnt"
	expectedIndex := 6

	outputMarker, ouputIndex := getMarker(input, 4)

	assert.Equal(t, expectedMarker, outputMarker, "should match")
	assert.Equal(t, expectedIndex, ouputIndex, "should match")
}

func Test_getMarker_sample5(t *testing.T) {
	input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	expectedMarker := "zqfr"
	expectedIndex := 7

	outputMarker, ouputIndex := getMarker(input, 4)

	assert.Equal(t, expectedMarker, outputMarker, "should match")
	assert.Equal(t, expectedIndex, ouputIndex, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
