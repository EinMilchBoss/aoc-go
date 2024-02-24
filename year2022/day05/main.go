package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

func parseSeeds(line string) (seeds []int) {
	for _, value := range strings.Split(line, " ")[1:] {
		converted, _ := strconv.Atoi(value)
		seeds = append(seeds, converted)
	}
	return
}

type almanacRange struct {
	dst int
	src int
	len int
}

type almanacMap []almanacRange

type almanac []almanacMap

func parseMap(block string) almanacMap {
	var values [3]int

	lines := strings.Split(block, "\n")[1:]
	ranges := make([]almanacRange, 0, len(lines))
	for _, line := range lines {
		for i, value := range strings.Split(line, " ") {
			converted, _ := strconv.Atoi(value)
			values[i] = converted
		}

		ranges = append(ranges, almanacRange{
			dst: values[0],
			src: values[1],
			len: values[2],
		})
	}

	return ranges
}

func parseAlmanac(blocks [7]string) almanac {
	maps := make([]almanacMap, 0, 7)
	for _, block := range blocks {
		maps = append(maps, parseMap(block))
	}
	return maps
}

func convertByMap(value int, amap almanacMap) int {
	for _, ar := range amap {
		if ar.src <= value && value < ar.src+ar.len {
			offset := value - ar.src
			return ar.dst + offset
		}
	}
	return value
}

func toLocation(seed int, a almanac) int {
	value := seed
	for _, am := range a {
		value = convertByMap(value, am)
	}
	return value
}

func partOne(data []byte) int {
	// parse seeds
	// parse almanac
	// -> parse maps
	//   -> parse ranges

	blocks := strings.Split(string(data), "\n\n")
	seeds := parseSeeds(blocks[0])
	almanac := parseAlmanac([7]string(blocks[1:]))

	for i, s := range seeds {
		seeds[i] = toLocation(s, almanac)
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
