package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

type OppMove rune

const (
	OppRock     OppMove = 'A'
	OppPaper    OppMove = 'B'
	OppScissors OppMove = 'C'
)

type PlyMove rune

const (
	PlyRock     PlyMove = 'X'
	PlyPaper    PlyMove = 'Y'
	PlyScissors PlyMove = 'Z'
)

type Outcome uint

const (
	Defeat = Outcome(iota * 3)
	Draw
	Victory
)

func partOne(data []byte) uint {
	input := string(data)

	var points uint
	for _, line := range strings.Split(input, "\n") {
		runes := []rune(line)
		opp, ply := OppMove(runes[0]), PlyMove(runes[2])
		outcome, move := uint(getOutcome(opp, ply)), getMovePoints(ply)
		points += outcome + move
	}

	return points
}

func partTwo(data []byte) int {
	return 0
}

func getMovePoints(move PlyMove) uint {
	switch move {
	case PlyRock:
		return 1
	case PlyPaper:
		return 2
	default:
		return 3
	}
}

func getOutcome(opp OppMove, ply PlyMove) Outcome {
	switch opp {
	case OppRock:
		switch ply {
		case PlyRock:
			return Draw
		case PlyPaper:
			return Victory
		case PlyScissors:
			return Defeat
		}
	case OppPaper:
		switch ply {
		case PlyRock:
			return Defeat
		case PlyPaper:
			return Draw
		case PlyScissors:
			return Victory
		}
	case OppScissors:
		switch ply {
		case PlyRock:
			return Victory
		case PlyPaper:
			return Defeat
		case PlyScissors:
			return Draw
		}
	}
	panic("Reached non-deterministic outcome.")
}

func doesWin(opp OppMove, ply PlyMove) bool {
	return (opp == OppRock && ply == PlyPaper) ||
		(opp == OppPaper && ply == PlyScissors) ||
		(opp == OppScissors && ply == PlyRock)
}

func main() {
	inputs, err := aoc.ReadInputs(2022, 2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input files could not be read.\nError: %s\n", err)
		os.Exit(1)
	}
	one := aoc.PartOne(&inputs, partOne)
	two := aoc.PartTwo(&inputs, partTwo)

	fmt.Printf("%s\n\n", one.GetTestProtocol(15))
	fmt.Printf("%s\n\n", two.GetTestProtocol(12))

	fmt.Printf("Result part one:\n%v\n", one.Run())
	//fmt.Printf("Result part two:\n%v\n", two.Run())
}
