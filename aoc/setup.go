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

func ReadInputs(year, day uint) (inputs Inputs, err error) {
	var example, actual []byte
	example, err = readInput(year, day, exampleInputFilename)
	if err != nil {
		return
	}
	actual, err = readInput(year, day, actualInputFilename)
	if err != nil {
		return
	}

	inputs = Inputs{
		Example: example,
		Actual:  actual,
	}
	return
}

func readInput(year, day uint, filename string) (data []byte, err error) {
	file := fmt.Sprintf("./res/year%04v/day%02v/%s", year, day, filename)
	data, err = os.ReadFile(file)
	return
}
