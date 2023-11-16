package aoc

import (
	"fmt"
	"strings"
)

type Solution[T comparable] func([]byte) T

type Part[T comparable] struct {
	number   uint
	solution Solution[T]
	inputs   *Inputs
}

func PartOne[T comparable](inputs *Inputs, solution Solution[T]) Part[T] {
	return Part[T]{
		number:   1,
		solution: solution,
		inputs:   inputs,
	}
}

func PartTwo[T comparable](inputs *Inputs, solution Solution[T]) Part[T] {
	return Part[T]{
		number:   2,
		solution: solution,
		inputs:   inputs,
	}
}

func (p Part[T]) GetTestProtocol(expected T) string {
	solutionResult := p.solution(p.inputs.Example)

	var testResult string
	if solutionResult == expected {
		testResult = "PASS"
	} else {
		testResult = "FAIL"
	}

	lines := make([]string, 0, 5)
	lines = append(lines, fmt.Sprintf("--- PART %d ---", p.number))
	lines = append(lines, fmt.Sprintf("Expected:\n%v", expected))
	lines = append(lines, fmt.Sprintf("Actual:\n%v", solutionResult))
	lines = append(lines, fmt.Sprintf("---- %s ----", testResult))
	return strings.Join(lines, "\n")
}

func (p Part[T]) Run() T {
	return p.solution(p.inputs.Actual)
}
