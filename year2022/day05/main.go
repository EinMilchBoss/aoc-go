package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
	"github.com/einmilchboss/aoc-go/year2022/day05/almanac"
)

func partOne(data []byte) int {
	blocks := strings.Split(string(data), "\n\n")
	seeds := almanac.ParseSeeds(blocks[0])
	almanac := almanac.ParseAlmanac([7]string(blocks[1:]))

	for i, s := range seeds {
		seeds[i] = almanac.Location(s)
	}

	return slices.Min(seeds)
}

func partTwo(data []byte) int {
	return 0
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	//two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol(35))
	//fmt.Printf("%s\n\n", two.GetTestProtocol(0))

	fmt.Printf("Result part one:\n%v\n", one.Run())
	//fmt.Printf("Result part two:\n%v\n", two.Run())
}
