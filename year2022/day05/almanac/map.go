package almanac

import (
	"strconv"
	"strings"
)

type AlmanacMap []AlmanacRange

func ParseAlmanacMap(block string) AlmanacMap {
	var values [3]int

	lines := strings.Split(block, "\n")[1:]
	ranges := make([]AlmanacRange, len(lines))
	for i, line := range lines {
		for j, value := range strings.Split(line, " ") {
			converted, _ := strconv.Atoi(value)
			values[j] = converted
		}

		ranges[i] = AlmanacRange{
			dst: values[0],
			src: values[1],
			len: values[2],
		}
	}

	return ranges
}

func (am AlmanacMap) convert(value int) int {
	for _, ar := range am {
		if ar.src <= value && value < ar.src+ar.len {
			offset := value - ar.src
			return ar.dst + offset
		}
	}
	return value
}
