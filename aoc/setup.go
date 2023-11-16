package aoc

import (
	"fmt"
	"os"
)

const (
	exampleInputFilename = "example.txt"
	actualInputFilename  = "actual.txt"
)

type Inputs struct {
	Example []byte
	Actual  []byte
}

func ReadInputs(year, day uint) Inputs {
	return Inputs{
		Example: readInput(year, day, exampleInputFilename),
		Actual:  readInput(year, day, actualInputFilename),
	}
}

func readInput(year, day uint, filename string) []byte {
	file := fmt.Sprintf("../../res/year%04v/day%02v/%v", year, day, filename)
	data, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Something went wrong reading input file %s. %s", file, err))
	}
	return data
}
