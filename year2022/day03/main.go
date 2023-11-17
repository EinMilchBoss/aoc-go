package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

func partOne(data []byte) (sum uint) {
	input := string(data)
	for _, line := range strings.Split(input, "\n") {
		half := len(line) / 2
		left, right := line[:half], line[half:]
		for _, r := range left {
			if strings.ContainsRune(right, r) {
				sum += priority(r)
				break
			}
		}
	}

	return sum
}

func partTwo(data []byte) (sum uint) {
	input := string(data)
	lines := strings.Split(input, "\n")
	for _, lines := range chunk(lines, 3) {
		for _, r := range lines[0] {
			if strings.ContainsRune(lines[1], r) && strings.ContainsRune(lines[2], r) {
				sum += priority(r)
				break
			}
		}
	}
	return
}

func chunk(lines []string, size int) [][]string {
	amount := len(lines) / size
	chunks := make([][]string, 0, amount)
	for i := 0; i < amount; i += 1 {
		chunk := make([]string, 0, size)
		for j := 0; j < size; j += 1 {
			chunk = append(chunk, lines[i*size+j])
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func priority(r rune) uint {
	switch {
	case 'a' <= r && r <= 'z':
		return uint(r-'a') + 1
	case 'A' <= r && r <= 'Z':
		return uint(r-'A') + 1 + 26
	}
	panic("Encountered illegal rune.")
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol(157))
	fmt.Printf("%s\n\n", two.GetTestProtocol(70))

	fmt.Printf("Result part one:\n%v\n", one.Run())
	fmt.Printf("Result part two:\n%v\n", two.Run())
}
