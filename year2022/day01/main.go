package main

import (
	"fmt"
	"os"

	"github.com/einmilchboss/aoc-go/aoc"
)

func partOne(input []byte) string {
	return "Not defined."
}

func partTwo(input []byte) string {
	return "Not defined."
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol("24000"))
	fmt.Printf("%s\n\n", two.GetTestProtocol("45000"))

	fmt.Printf("Result part one:\n%s\n", one.Run())
	fmt.Printf("Result part two:\n%s\n", two.Run())
}
