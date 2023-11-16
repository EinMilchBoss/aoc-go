package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/einmilchboss/aoc-go/aoc"
)

type first rune

const (
	fstRock     first = 'A'
	fstPaper    first = 'B'
	fstScissors first = 'C'
)

type second rune

const (
	sndRock     second = 'X'
	sndPaper    second = 'Y'
	sndScissors second = 'Z'
)

type outcome uint

const (
	defeat = outcome(iota * 3)
	draw
	victory
)

func partOne(data []byte) uint {
	input := string(data)

	var points uint
	for _, line := range strings.Split(input, "\n") {
		runes := []rune(line)
		opp, ply := first(runes[0]), second(runes[2])
		out := deduceOutcome(opp, ply)

		outPoints := uint(out)
		movePoints := pointsFromSecond(ply)
		points += outPoints + movePoints
	}

	return points
}

func partTwo(data []byte) uint {
	input := string(data)

	var points uint
	for _, line := range strings.Split(input, "\n") {
		runes := []rune(line)
		opp, out := first(runes[0]), secondToOutcome(second(runes[2]))
		ply := deducePlayerMove(opp, out)

		outPoints := uint(out)
		movePoints := pointsFromSecond(ply)
		points += outPoints + movePoints
	}

	return points
}

func deduceOutcome(opp first, ply second) outcome {
	switch opp {
	case fstRock:
		switch ply {
		case sndRock:
			return draw
		case sndPaper:
			return victory
		case sndScissors:
			return defeat
		}
	case fstPaper:
		switch ply {
		case sndRock:
			return defeat
		case sndPaper:
			return draw
		case sndScissors:
			return victory
		}
	case fstScissors:
		switch ply {
		case sndRock:
			return victory
		case sndPaper:
			return defeat
		case sndScissors:
			return draw
		}
	}
	panic("Played non-deterministic outcome.")
}

func pointsFromSecond(snd second) uint {
	switch snd {
	case sndRock:
		return 1
	case sndPaper:
		return 2
	case sndScissors:
		return 3
	}
	panic("Used illegal move.")
}

func secondToOutcome(snd second) outcome {
	switch snd {
	case sndRock:
		return defeat
	case sndPaper:
		return draw
	case sndScissors:
		return victory
	}
	panic("Played non-deterministic outcome.")
}

func deducePlayerMove(fst first, out outcome) second {
	switch fst {
	case fstRock:
		switch out {
		case defeat:
			return sndScissors
		case draw:
			return sndRock
		case victory:
			return sndPaper
		}
	case fstPaper:
		switch out {
		case defeat:
			return sndRock
		case draw:
			return sndPaper
		case victory:
			return sndScissors
		}
	case fstScissors:
		switch out {
		case defeat:
			return sndPaper
		case draw:
			return sndScissors
		case victory:
			return sndRock
		}
	}
	panic("Used illegal move.")
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
	fmt.Printf("Result part two:\n%v\n", two.Run())
}
