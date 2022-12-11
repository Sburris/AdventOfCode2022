package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DirectoryStack struct {
	list []string
}

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	CreateFileSystem()

	input := getInput(inputFile)
	fs := ProcessInput(input)

	dirList := fs.RootFolder.GetList("/")

	sum := 0
	for _, size := range dirList {
		if size <= 100000 {
			sum += size
		}
	}

	fmt.Print(sum)
}

type FileSystem struct {
	CurrentPath DirectoryStack
	CurrentFolder *DirectoryData
	RootFolder DirectoryData
}

type FileData struct {
	Size int
	Name string
}

type DirectoryData struct {
	Name string
	Subdirectories map[string]*DirectoryData
	ParentDirectory *DirectoryData
	Files []*FileData
}

func ProcessInput(input []string) FileSystem {
	fs := CreateFileSystem()

	for _, line := range input {
		fs.ProcessInput(line)
	}

	return fs
}

func CreateFileSystem() FileSystem {
	rootDir := CreateDirectory("root", nil)

	return FileSystem{
		CurrentPath: CreateStack(),
		CurrentFolder: &rootDir,
		RootFolder: rootDir,
	}
}

func (dir *DirectoryData) GetList(prefix string) map[string]int {
	mySize := dir.GetSize()
	myPath := prefix + "/" + dir.Name
	if dir.ParentDirectory == nil {
		myPath = "/"
	}

	list := make(map[string]int)
	list[myPath] = mySize

	for _, subDir := range dir.Subdirectories {
		subList := subDir.GetList(myPath)

		for key, value := range subList {
			list[key] = value
		}
	}

	return list
}

func (dir *DirectoryData) GetSize() int {
	sum := 0

	for _, file := range dir.Files {
		sum += file.Size
	}

	for _, subDirs := range dir.Subdirectories {
		sum += subDirs.GetSize()
	}

	return sum
}

func (fs *FileSystem) ProcessInput(input string) {
	if fs.isCommand(input) == true {
		fs.processCommand(input)
	} else {
		fs.processContent(input)

	}
}

func (fs *FileSystem) processContent(input string) {
	parts := strings.Split(input, " ")
	if parts[0] == "dir" {
		fs.CurrentFolder.CreateSubdirectory(parts[1])
	} else {
		name := parts[1]
		size, _ := strconv.Atoi(parts[0])
		file := CreateFileData(name, size)
		fs.CurrentFolder.AddFile(file)
	}
}

func (dir *DirectoryData) AddFile(file FileData) {
	dir.Files = append(dir.Files, &file)
}

func (fs *FileSystem) isCommand(input string) bool {
	return input[0] == '$'
}

func (fs *FileSystem) processCommand(input string) {
	parts := strings.Split(input, " ")
	switch parts[1] {
	case "cd":
		fs.changeDirectory(parts[2])
	case "ls":
		// Do nothing
	default:
		panic(fmt.Sprintf("UNKNOWN COMMAND: %s", input))
	}
}

func (fs *FileSystem) changeDirectory(dirName string) {
	if dirName == "/" {
		fs.CurrentPath.GoToRoot()
		fs.CurrentFolder = &fs.RootFolder
	} else if dirName == ".." {
		fs.CurrentPath.Pop()
		fs.CurrentFolder = fs.CurrentFolder.ParentDirectory
	} else {
		dir, found := fs.CurrentFolder.Subdirectories[dirName]
		if found == false {
			fs.CurrentFolder.CreateSubdirectory(dirName)
			dir = fs.CurrentFolder.Subdirectories[dirName]
		}
		fs.CurrentPath.Push(dirName)
		fs.CurrentFolder = dir
	}
}

func (dir *DirectoryData) CreateSubdirectory(name string) {
	newDir := CreateDirectory(name, dir)
	dir.Subdirectories[name] = &newDir
}

func CreateDirectory(name string, parentDir *DirectoryData) DirectoryData {
	return DirectoryData{
		Name: name,
		Subdirectories: make(map[string]*DirectoryData, 0),
		ParentDirectory: parentDir,
		Files: make([]*FileData, 0),
	}
}

func CreateFileData(name string, size int) FileData {
	return FileData{
		Name: name,
		Size: size,
	}
}

func isFileData(input string) (bool, FileData) {
	parts := strings.Split(input, " ")
	isNum := isNumber(parts[0])
	if isNum == true {
		filename := parts[1]
		fileSize, _ := strconv.Atoi(parts[0])
		fileData := CreateFileData(filename, fileSize)
		return true, fileData
	}

	return false, FileData{}
}

var numberPattern = "^[0-9]+$"
var numberRegex = regexp.MustCompile(numberPattern)

func isNumber(input string) bool {
	result := numberRegex.MatchString(input)
	return result
}

// Read in raw data so we can process it
func getInput(filename string) []string {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/// ---------------------------------------------------------------------------
///   Stack
/// ---------------------------------------------------------------------------

func CreateStack() DirectoryStack {
	return DirectoryStack{
		list: make([]string, 0),
	}
}

func (s *DirectoryStack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *DirectoryStack) Push(value string) {
	s.list = append(s.list, value)
}

func (s *DirectoryStack) Pop() (string, bool) {
	if s.IsEmpty() == true {
		return "", false
	}
	index := len(s.list) - 1
	element := s.list[index]
	s.list = s.list[:index]
	return element, true
}

func (s *DirectoryStack) Peek() string {
	return s.list[len(s.list)-1]
}

func (s *DirectoryStack) GoToRoot() {
	s.list = make([]string, 0)
}

func (s *DirectoryStack) GetPath() string {
	return "/" + strings.Join(s.list, "/")
}
