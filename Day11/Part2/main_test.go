package main

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   CreateMonkey
/// ---------------------------------------------------------------------------

func Test_CreateMonkey_sampleMonkey(t *testing.T) {
	input := []string{
		"  Starting items: 79, 98",
		"  Operation: new = old * 19",
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
	}
	expected := Monkey{
		Items:           []big.Int{*big.NewInt(79), *big.NewInt(98)},
		Operations:      []big.Int{*big.NewInt(19), *big.NewInt(0)},
		IsDoubled:       false,
		DivisibleBy:     *big.NewInt(23),
		Test:            []int{2, 3},
		InspectionCount: 0,
	}

	output := CreateMonkey(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
