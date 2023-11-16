package aoc

import (
	"fmt"
	"strings"
)

type Solution func([]byte) string

type Part struct {
	number   uint
	solution Solution
	inputs   *Inputs
}

func PartOne(inputs *Inputs, solution Solution) Part {
	return Part{
		number:   1,
		solution: solution,
		inputs:   inputs,
	}
}

func PartTwo(inputs *Inputs, solution Solution) Part {
	return Part{
		number:   2,
		solution: solution,
		inputs:   inputs,
	}
}

func (p Part) GetTestProtocol(expected string) string {
	solutionResult := p.solution(p.inputs.Example)

	var testResult string
	if solutionResult == expected {
		testResult = "PASS"
	} else {
		testResult = "FAIL"
	}

	lines := make([]string, 0, 5)
	lines = append(lines, fmt.Sprintf("--- PART %d ---", p.number))
	lines = append(lines, fmt.Sprintf("Expected:\n%s", expected))
	lines = append(lines, fmt.Sprintf("Actual:\n%s", solutionResult))
	lines = append(lines, fmt.Sprintf("---- %s ----", testResult))
	return strings.Join(lines, "\n")
}

func (p Part) Run() string {
	return p.solution(p.inputs.Actual)
}
