package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cpu struct {
	Register []int
}

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"

	input := getInput(inputFile)
	cpu := CreateCpu()

	for _, instruction := range input {
		cpu.ProcessInstruction(instruction)
	}

	keyValues := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, position := range keyValues {
		registerValue := cpu.Register[position-1]
		product := registerValue * position
		fmt.Printf("%d * %d = %d\n", position, registerValue, product)
		sum += (product)
	}

	fmt.Printf("Sum: %d\n", sum)
}

func CreateCpu() Cpu {
	return Cpu{
		Register: []int{1},
	}
}

func (cpu *Cpu) ProcessInstruction(instruction string) {
	parts := strings.Split(instruction, " ")
	lastRegisterValue := cpu.Register[len(cpu.Register)-1]
	switch parts[0] {
	case "noop":
		cpu.Register = append(cpu.Register, lastRegisterValue)
	case "addx":
		value, _ := strconv.Atoi(parts[1])
		cpu.Register = append(cpu.Register, lastRegisterValue)
		newValue := lastRegisterValue + value
		cpu.Register = append(cpu.Register, newValue)
	default:
		panic("UNKNOWN instruction")
	}
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
