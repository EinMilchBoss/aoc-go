package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

func partOne(data []byte) int {
	return sumElvesMaxCal(data, 1)
}

func partTwo(data []byte) int {
	return sumElvesMaxCal(data, 3)
}

func sumElvesMaxCal(data []byte, maxN int) int {
	input := string(data)

	maxElves := make([]int, maxN)
	for _, elf := range strings.Split(input, "\n\n") {
		cals := sumCals(elf)

		idxMin, valMin := findMin(maxElves[:])
		if cals > valMin {
			maxElves[idxMin] = cals
		}
	}

	return sum(maxElves[:])
}

func sumCals(elf string) int {
	cals := 0
	for _, line := range strings.Split(elf, "\n") {
		cal, _ := strconv.Atoi(line)
		cals += cal
	}
	return cals
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
