package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

type pair struct {
	from int
	to   int
}

func (left pair) includesAll(right pair) bool {
	return left.from >= right.from && right.to >= left.to
}

func (left pair) includesAny(right pair) bool {
	for l := left.from; l <= left.to; l++ {
		for r := right.from; r <= right.to; r++ {
			if l == r {
				return true
			}
		}
	}
	return false
}

func partOne(data []byte) (sum int) {
	//fmt.Println(string(data))

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		sides := strings.SplitN(line, ",", 2)
		left := parsePair(sides[0])
		right := parsePair(sides[1])
		if left.includesAll(right) || right.includesAll(left) {
			sum++
		}
	}

	return
}

func partTwo(data []byte) (sum int) {
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		sides := strings.SplitN(line, ",", 2)
		left := parsePair(sides[0])
		right := parsePair(sides[1])
		if left.includesAny(right) || right.includesAny(left) {
			sum++
		}
	}

	return
}

func parsePair(side string) pair {
	values := strings.SplitN(side, "-", 2)
	from, _ := strconv.Atoi(values[0])
	to, _ := strconv.Atoi(values[1])
	return pair{from, to}
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
