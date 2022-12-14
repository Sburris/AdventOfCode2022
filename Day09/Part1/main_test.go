package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   Move
/// ---------------------------------------------------------------------------

func Test_Move_StartUp(t *testing.T) {
	rope := CreateRope()

	rope.move("U")

	assert.Equal(t, 0, rope.Head.x, "should match")
	assert.Equal(t, 1, rope.Head.y, "should match")
	assert.Equal(t, 0, rope.Tail.x, "should match")
	assert.Equal(t, 0, rope.Tail.y, "should match")
}

func Test_Move_StartDown(t *testing.T) {
	rope := CreateRope()

	rope.move("D")

	assert.Equal(t, 0, rope.Head.x, "should match")
	assert.Equal(t, -1, rope.Head.y, "should match")
	assert.Equal(t, 0, rope.Tail.x, "should match")
	assert.Equal(t, 0, rope.Tail.y, "should match")
}

func Test_Move_StartLeft(t *testing.T) {
	rope := CreateRope()

	rope.move("L")

	assert.Equal(t, -1, rope.Head.x, "should match")
	assert.Equal(t, 0, rope.Head.y, "should match")
	assert.Equal(t, 0, rope.Tail.x, "should match")
	assert.Equal(t, 0, rope.Tail.y, "should match")
}

func Test_Move_StartRight(t *testing.T) {
	rope := CreateRope()

	rope.move("R")

	assert.Equal(t, 1, rope.Head.x, "should match")
	assert.Equal(t, 0, rope.Head.y, "should match")
	assert.Equal(t, 0, rope.Tail.x, "should match")
	assert.Equal(t, 0, rope.Tail.y, "should match")
}

func Test_Move_Exmaple_Step3(t *testing.T) {
	rope := CreateRope()
	rope.Head = RopeEnd{x: 2, y: 0}
	rope.Tail = RopeEnd{x: 1, y: 0}

	rope.move("R")

	assert.Equal(t, 3, rope.Head.x, "should match")
	assert.Equal(t, 0, rope.Head.y, "should match")
	assert.Equal(t, 2, rope.Tail.x, "should match")
	assert.Equal(t, 0, rope.Tail.y, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
