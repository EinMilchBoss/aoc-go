package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

func partOne(data []byte) (maxCals int) {
	input := string(data)

	var cals int
	for _, elf := range strings.Split(input, "\n\n") {
		cals = 0
		for _, line := range strings.Split(elf, "\n") {
			cal, _ := strconv.Atoi(line)
			cals += cal
		}

		if cals > maxCals {
			maxCals = cals
		}
	}

	return
}

func partTwo(data []byte) int {
	input := string(data)

	var topElves [3]int
	var cals int
	for _, elf := range strings.Split(input, "\n\n") {
		cals = 0
		for _, line := range strings.Split(elf, "\n") {
			cal, _ := strconv.Atoi(line)
			cals += cal
		}

		idxMin, valMin := findMin(topElves[:])
		if cals > valMin {
			topElves[idxMin] = cals
		}
	}

	return sum(topElves[:])
}

func findMin(nums []int) (idxMin int, valMin int) {
	valMin = nums[0]
	for i := 1; i < len(nums); i += 1 {
		if nums[i] < valMin {
			idxMin = i
			valMin = nums[i]
		}
	}
	return
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol(24000))
	fmt.Printf("%s\n\n", two.GetTestProtocol(45000))

	fmt.Printf("Result part one:\n%v\n", one.Run())
	fmt.Printf("Result part two:\n%v\n", two.Run())
}
