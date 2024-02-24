package almanac

import (
	"strconv"
	"strings"
)

type almanac [7]AlmanacMap

func ParseAlmanac(blocks [7]string) almanac {
	maps := [7]AlmanacMap{}
	for i, block := range blocks {
		maps[i] = ParseAlmanacMap(block)
	}
	return maps
}

func ParseSeeds(line string) []int {
	values := strings.Split(line, " ")[1:]
	seeds := make([]int, len(values))
	for i, value := range values {
		converted, _ := strconv.Atoi(value)
		seeds[i] = converted
	}
	return seeds
}

func (a almanac) Location(seed int) int {
	value := seed
	for _, am := range a {
		value = am.convert(value)
	}
	return value
}
