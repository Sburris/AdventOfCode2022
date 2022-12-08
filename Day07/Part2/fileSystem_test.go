package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   CreateFileSystem
/// ---------------------------------------------------------------------------

func Test_CreateFileSystem_default(t *testing.T) {
	rootDir := CreateDirectory("root", nil)
	expected := FileSystem {
		CurrentPath: CreateStack(),
		CurrentFolder: &rootDir,
		RootFolder: rootDir,
	}

	output := CreateFileSystem()

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   ProcessInput
/// ---------------------------------------------------------------------------

func Test_ProcessInput_twosubdirsoneperlevel(t *testing.T) {
	input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"$ cd a",
		"$ ls",
		"dir e",
	}
	fs := CreateFileSystem()
	for _, line := range input {
		fs.ProcessInput(line)
	}

	assert.Equal(t, 1, len(fs.RootFolder.Subdirectories), "should match")
	assert.Equal(t, 1, len(fs.RootFolder.Subdirectories["a"].Subdirectories), "should match")
}

func Test_ProcessInput_Files(t *testing.T) {
	input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"1234 root.txt",
		"$ cd a",
		"$ ls",
		"dir e",
		"1234 a.txt",
	}
	fs := CreateFileSystem()
	for _, line := range input {
		fs.ProcessInput(line)
	}

	assert.Equal(t, 1, len(fs.RootFolder.Files), "should match")
	assert.Equal(t, 1, len(fs.RootFolder.Subdirectories["a"].Files), "should match")
}
