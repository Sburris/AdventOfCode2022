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

	assert.Equal(t, 0, rope.Knots[0].x, "should match")
	assert.Equal(t, 1, rope.Knots[0].y, "should match")
	assert.Equal(t, 0, rope.Knots[1].x, "should match")
	assert.Equal(t, 0, rope.Knots[1].y, "should match")
}

func Test_Move_StartDown(t *testing.T) {
	rope := CreateRope()

	rope.move("D")

	assert.Equal(t, 0, rope.Knots[0].x, "should match")
	assert.Equal(t, -1, rope.Knots[0].y, "should match")
	assert.Equal(t, 0, rope.Knots[1].x, "should match")
	assert.Equal(t, 0, rope.Knots[1].y, "should match")
}

func Test_Move_StartLeft(t *testing.T) {
	rope := CreateRope()

	rope.move("L")

	assert.Equal(t, -1, rope.Knots[0].x, "should match")
	assert.Equal(t, 0, rope.Knots[0].y, "should match")
	assert.Equal(t, 0, rope.Knots[1].x, "should match")
	assert.Equal(t, 0, rope.Knots[1].y, "should match")
}

func Test_Move_StartRight(t *testing.T) {
	rope := CreateRope()

	rope.move("R")

	assert.Equal(t, 1, rope.Knots[0].x, "should match")
	assert.Equal(t, 0, rope.Knots[0].y, "should match")
	assert.Equal(t, 0, rope.Knots[1].x, "should match")
	assert.Equal(t, 0, rope.Knots[1].y, "should match")
}

/// ---------------------------------------------------------------------------
///   MoveInstruction
/// ---------------------------------------------------------------------------

func Test_MoveInstruction_FirstMove(t *testing.T) {
	rope := CreateRope()

	rope.MoveInstruction("R 5")

	assert.Equal(t, 5, rope.Knots[0].x, "should match")
	assert.Equal(t, 0, rope.Knots[0].y, "should match")
	assert.Equal(t, 4, rope.Knots[1].x, "should match")
	assert.Equal(t, 0, rope.Knots[1].y, "should match")
	assert.Equal(t, 3, rope.Knots[2].x, "should match")
	assert.Equal(t, 0, rope.Knots[2].y, "should match")
	assert.Equal(t, 2, rope.Knots[3].x, "should match")
	assert.Equal(t, 0, rope.Knots[3].y, "should match")
	assert.Equal(t, 1, rope.Knots[4].x, "should match")
	assert.Equal(t, 0, rope.Knots[4].y, "should match")
	assert.Equal(t, 0, rope.Knots[5].x, "should match")
	assert.Equal(t, 0, rope.Knots[5].y, "should match")
	assert.Equal(t, 0, rope.Knots[6].x, "should match")
	assert.Equal(t, 0, rope.Knots[6].y, "should match")
	assert.Equal(t, 0, rope.Knots[7].x, "should match")
	assert.Equal(t, 0, rope.Knots[7].y, "should match")
	assert.Equal(t, 0, rope.Knots[8].x, "should match")
	assert.Equal(t, 0, rope.Knots[8].y, "should match")
	assert.Equal(t, 0, rope.Knots[9].x, "should match")
	assert.Equal(t, 0, rope.Knots[9].y, "should match")
}

func Test_MoveInstruction_SecondMove(t *testing.T) {
	rope := CreateRope()

	rope.MoveInstruction("R 5")
	rope.MoveInstruction("U 8")

	assert.Equal(t, 5, rope.Knots[0].x, "should match") // H (5, 8)
	assert.Equal(t, 8, rope.Knots[0].y, "should match")
	assert.Equal(t, 5, rope.Knots[1].x, "should match") // 1 (5, 7)
	assert.Equal(t, 7, rope.Knots[1].y, "should match")
	assert.Equal(t, 5, rope.Knots[2].x, "should match") // 2 (5, 6)
	assert.Equal(t, 6, rope.Knots[2].y, "should match")
	assert.Equal(t, 5, rope.Knots[3].x, "should match") // 3 (5, 5)
	assert.Equal(t, 5, rope.Knots[3].y, "should match")
	assert.Equal(t, 5, rope.Knots[4].x, "should match") // 4 (5, 4)
	assert.Equal(t, 4, rope.Knots[4].y, "should match")
	assert.Equal(t, 4, rope.Knots[5].x, "should match") // 5 (4,4)
	assert.Equal(t, 4, rope.Knots[5].y, "should match")
	assert.Equal(t, 3, rope.Knots[6].x, "should match") // 6 (3,3)
	assert.Equal(t, 3, rope.Knots[6].y, "should match")
	assert.Equal(t, 2, rope.Knots[7].x, "should match") // 7 (2,2)
	assert.Equal(t, 2, rope.Knots[7].y, "should match")
	assert.Equal(t, 1, rope.Knots[8].x, "should match") // 8 (1,1)
	assert.Equal(t, 1, rope.Knots[8].y, "should match")
	assert.Equal(t, 0, rope.Knots[9].x, "should match") // 9 (0,0)
	assert.Equal(t, 0, rope.Knots[9].y, "should match")
}

func Test_MoveInstruction_ThirdMove(t *testing.T) {
	rope := CreateRope()

	rope.MoveInstruction("R 5")
	rope.MoveInstruction("U 8")
	rope.MoveInstruction("L 8")

	assert.Equal(t, -3, rope.Knots[0].x, "should match") // H (5, 8)
	assert.Equal(t, 8, rope.Knots[0].y, "should match")
	assert.Equal(t, -2, rope.Knots[1].x, "should match") // 1 (5, 7)
	assert.Equal(t, 8, rope.Knots[1].y, "should match")
	assert.Equal(t, -1, rope.Knots[2].x, "should match") // 2 (5, 6)
	assert.Equal(t, 8, rope.Knots[2].y, "should match")
	assert.Equal(t, 0, rope.Knots[3].x, "should match") // 3 (5, 5)
	assert.Equal(t, 8, rope.Knots[3].y, "should match")
	assert.Equal(t, 1, rope.Knots[4].x, "should match") // 4 (5, 4)
	assert.Equal(t, 8, rope.Knots[4].y, "should match")
	assert.Equal(t, 1, rope.Knots[5].x, "should match") // 5 (4,4)
	assert.Equal(t, 7, rope.Knots[5].y, "should match")
	assert.Equal(t, 1, rope.Knots[6].x, "should match") // 6 (3,3)
	assert.Equal(t, 6, rope.Knots[6].y, "should match")
	assert.Equal(t, 1, rope.Knots[7].x, "should match") // 7 (2,2)
	assert.Equal(t, 5, rope.Knots[7].y, "should match")
	assert.Equal(t, 1, rope.Knots[8].x, "should match") // 8 (1,1)
	assert.Equal(t, 4, rope.Knots[8].y, "should match")
	assert.Equal(t, 1, rope.Knots[9].x, "should match") // 9 (0,0)
	assert.Equal(t, 3, rope.Knots[9].y, "should match")

	assert.True(t, rope.TailHistory["0,0"], "should be true")
	assert.True(t, rope.TailHistory["1,1"], "should be true")
	assert.True(t, rope.TailHistory["2,2"], "should be true")
	assert.True(t, rope.TailHistory["1,3"], "should be true")
	assert.Equal(t, 4, len(rope.TailHistory), "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getOverlapWorkordersCount_sampleInput(t *testing.T) {
	input := TEST_INPUT_FILE
	expected := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "should match")
}
