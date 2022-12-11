package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// const TEST_INPUT_FILE string = "test_data/sample_input.txt"

// /// ---------------------------------------------------------------------------
// ///   CreateRucksack
// /// ---------------------------------------------------------------------------

// func Test_CreateRucksack_example1(t *testing.T) {
// 	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
// 	expected := Rucksack{
// 		compartmentOne: "vJrwpWtwJgWr",
// 		compartmentTwo: "hcsFMMfFFhFp",
// 	}

// 	output := CreateRucksack(input)

// 	assert.Equal(t, expected, output, "should match")
// }

// /// ---------------------------------------------------------------------------
// ///   GetSupplyItemPriority
// /// ---------------------------------------------------------------------------

// func Test_GetSupplyItemPriority_lowercaseA(t *testing.T) {
// 	input := "a"
// 	expected := 1

// 	output := GetSupplyItemPriority(input)

// 	assert.Equal(t, expected, output, "should match")
// }

// func Test_GetSupplyItemPriority_UppercaseA(t *testing.T) {
// 	input := "A"
// 	expected := 27

// 	output := GetSupplyItemPriority(input)

// 	assert.Equal(t, expected, output, "should match")
// }

// /// ---------------------------------------------------------------------------
// ///   Rucksack.GetOverlappingItem
// /// ---------------------------------------------------------------------------

// func Test_RucksackGetOverlappingItem_example1(t *testing.T) {
// 	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
// 	expected := "p"

// 	rucksack := CreateRucksack(input)
// 	output := rucksack.GetOverlappingItem()

// 	assert.Equal(t, expected, output, "should match")
// }

// /// ---------------------------------------------------------------------------
// ///   Rucksack.getDistinctItems
// /// ---------------------------------------------------------------------------

// // func Test_RucksackGetDistinctItems_example1(t *testing.T) {
// // 	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
// // 	expected := "vJrwpWtg"

// // 	rucksack := CreateRucksack(input)
// // 	output := rucksack.getDistinctItems(rucksack.compartmentOne)

// // 	assert.Equal(t, expected, output, "should match")
// // }


// /// ---------------------------------------------------------------------------
// ///   getInput
// /// ---------------------------------------------------------------------------

// func Test_getInput_sampleInput(t *testing.T) {
// 	input := TEST_INPUT_FILE
// 	expected := []string{
// 		"vJrwpWtwJgWrhcsFMMfFFhFp",
// 		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
// 		"PmmdzqPrVvPwwTWBwg",
// 		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
// 		"ttgJtRGJQctTZtZT",
// 		"CrZsJsPPZsGzwwsLwLmpwMDw",
// 	}

// 	output := getInput(input)

// 	assert.Equal(t, expected, output, "should match")
// }
