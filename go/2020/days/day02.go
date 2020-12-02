package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2020/util"
)

// Day02 - Password Philosophy
func Day02() {
	fmt.Println("-------------------------")
	fmt.Println("DAY TWO")

	input, err := util.FileAsStrings("./input/day02input.txt")
	if err != nil {
		panic(err)
	}

	part1Count := 0
	part2Count := 0
	for _, v := range input {
		part1Valid, part2Valid := validate(v)

		if part1Valid {
			part1Count++
		}

		if part2Valid {
			part2Count++
		}
	}

	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)

	fmt.Println("-------------------------")
	fmt.Println()
}

func validate(row string) (part1Valid bool, part2Valid bool) {
	parts := strings.Split(row, ":")
	rule := strings.Split(strings.TrimSpace(parts[0]), " ")
	digits := strings.Split(rule[0], "-")
	character := rule[1]

	from, err := strconv.Atoi(digits[0])
	if err != nil {
		return false, false
	}

	to, err := strconv.Atoi(digits[1])
	if err != nil {
		return false, false
	}

	value := strings.TrimSpace(parts[1])
	occurances := strings.Count(value, character)

	part1 := occurances >= from && occurances <= to
	part2 := (len(value) >= from && string(value[from-1]) == character) != (len(value) >= to && string(value[to-1]) == character)

	return part1, part2
}
