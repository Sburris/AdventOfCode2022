package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   getOverlapWorkordersCount
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := getInput(TEST_INPUT_FILE)
	expected := 4

	output := getOverlapWorkordersCount(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getWorkOrders
/// ---------------------------------------------------------------------------

func Test_getWorkOrders_example1(t *testing.T) {
	input := "2-4,6-8"
	eFirstStart := 2
	eFirstEnd := 4
	eSecondStart := 6
	eSecondEnd := 8

	oFirstStart, oFirstEnd, oSecondStart, oSecondEnd := getWorkOrders(input)

	assert.Equal(t, eFirstStart, oFirstStart, "should match")
	assert.Equal(t, eFirstEnd, oFirstEnd, "should match")
	assert.Equal(t, eSecondStart, oSecondStart, "should match")
	assert.Equal(t, eSecondEnd, oSecondEnd, "should match")
}

/// ---------------------------------------------------------------------------
///   doesContainOverlap
/// ---------------------------------------------------------------------------

func Test_doesContainOverlap_noOverlap(t *testing.T) {
	input := "2-4,6-8"
	expected := false

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainOverlap(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainOverlap_partialOverlap(t *testing.T) {
	input := "2-4,3-8"
	expected := true

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainOverlap(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainOverlap_fullOverlap(t *testing.T) {
	input := "2-4,1-5"
	expected := true

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainOverlap(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}



/// ---------------------------------------------------------------------------
///   doesContainSubset
/// ---------------------------------------------------------------------------

func Test_doesContainSubset_noOverlap(t *testing.T) {
	input := "2-4,6-8"
	expected := false

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainSubset_partialOverlapBeginning(t *testing.T) {
	input := "2-4,3-8"
	expected := false

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainSubset_partialOverlap(t *testing.T) {
	input := "2-4,3-8"
	expected := false

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainSubset_mirrored(t *testing.T) {
	input := "2-4,2-4"
	expected := true

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainSubset_firstOverlapsSecond(t *testing.T) {
	input := "1-5,2-4"
	expected := true

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}

func Test_doesContainSubset_secondOverlapsFirst(t *testing.T) {
	input := "2-4,1-5"
	expected := true

	start1, end1, start2, end2 := getWorkOrders(input)
	output := doesContainSubset(start1, end1, start2, end2)

	assert.Equal(t, expected, output, "should match")
}


/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getInput_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
