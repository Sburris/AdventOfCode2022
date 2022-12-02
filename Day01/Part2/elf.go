package main

import (
	"strconv"
)

type Elf struct {
	CaloriesCarried int
	Food []int
}

func CreateElf() Elf {
	return Elf{
		CaloriesCarried: 0,
		Food: make([]int, 0),
	}
}

func CreateElfWithFood(food []string) Elf {
	newElf := CreateElf()

	for _, foodItem := range food {
		newElf.AddFood(foodItem)
	}

	return newElf
}

func (elf *Elf) AddFood(food string) {
	calories, err := strconv.Atoi(food)
	checkErr(err)

	elf.CaloriesCarried += calories

	elf.Food = append(elf.Food, calories)
}
