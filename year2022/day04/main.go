package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

func partOne(data []byte) int {
	return sumIf(data, func(left, right pair) bool {
		return left.includesAll(right) || right.includesAll(left)
	})
}

func partTwo(data []byte) int {
	return sumIf(data, func(left, right pair) bool {
		return left.includesAny(right) || right.includesAny(left)
	})
}

func sumIf(data []byte, predicate func(left, right pair) bool) (sum int) {
	for _, line := range strings.Split(string(data), "\n") {
		left, right, err := parseLine(line)
		if err != nil {
			panic(fmt.Errorf(
				`encountered an issue for parsing line "%v" due to %w`,
				line,
				err,
			))
		}
		if predicate(left, right) {
			sum++
		}
	}
	return
}

func parseLine(line string) (left, right pair, err error) {
	sides := strings.SplitN(line, ",", 2)
	if left, err = parsePair(sides[0]); err != nil {
		return
	}
	right, err = parsePair(sides[1])
	return
}

func parsePair(side string) (p pair, err error) {
	values := strings.SplitN(side, "-", 2)
	if p.from, err = strconv.Atoi(values[0]); err != nil {
		return pair{}, err
	}
	if p.to, err = strconv.Atoi(values[1]); err != nil {
		return pair{}, err
	}
	return
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol(2))
	fmt.Printf("%s\n\n", two.GetTestProtocol(4))

	fmt.Printf("Result part one:\n%v\n", one.Run())
	fmt.Printf("Result part two:\n%v\n", two.Run())
}
