package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   IsEmpty
/// ---------------------------------------------------------------------------

func Test_stackIsEmpty_emptyStack(t *testing.T) {
	expected := true
	stack := CreateStack()

	output := stack.IsEmpty()

	assert.Equal(t, expected, output, "should match")
}

func Test_stackIsEmpty_oneItem(t *testing.T) {
	expected := false
	stack := CreateStack()

	stack.Push("a")
	output := stack.IsEmpty()

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   Push/Pop
/// ---------------------------------------------------------------------------

func Test_stackPushPop_oneItem(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "c"
	stack := CreateStack()

	for _, item := range input {
		stack.Push(item)
	}

	output, _ := stack.Pop(1)

	assert.Equal(t, expected, output, "should match")
}

func Test_stackPushPop_twoItems(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "bc"
	stack := CreateStack()

	for _, item := range input {
		stack.Push(item)
	}

	output, _ := stack.Pop(2)

	assert.Equal(t, expected, output, "should match")
}

