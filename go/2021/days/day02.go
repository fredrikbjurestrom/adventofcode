package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day02 - Dive!
func Day02() {
	fmt.Println("-------------------------")
	fmt.Println("DAY TWO")

	instructions, err := util.FileAsStrings("./input/day02input.txt")
	if err != nil {
		panic(err)
	}

	part1, err := getSubmarinePosition(instructions)
	if err != nil {
		panic(err)
	}

	part2, err := getSubmarinePositionWithAim(instructions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func getSubmarinePosition(instructions []string) (result int, err error) {
	horizontal := 0
	depth := 0

	for _, s := range instructions {
		parts := strings.Fields(s)

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	return horizontal * depth, nil
}

func getSubmarinePositionWithAim(instructions []string) (result int, err error) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, s := range instructions {
		parts := strings.Fields(s)

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	return horizontal * depth, nil
}
