package days

import (
	"errors"
	"fmt"

	"github.com/fredrikbjurestrom/adventofcode/go/2020/util"
)

// Day01 - Report Repair
func Day01() {
	fmt.Println("-------------------------")
	fmt.Println("DAY ONE")

	ints, err := util.FileAsInts("./input/day01input.txt")
	if err != nil {
		panic(err)
	}

	part1, part1Operation, err := findAddends(ints, 2, 2020)
	if err != nil {
		panic(err)
	}

	part2, part2Operation, err := findAddends(ints, 3, 2020)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1, "(", part1Operation, ")")
	fmt.Println("Part 2:", part2, "(", part2Operation, ")")

	fmt.Println("-------------------------")
	fmt.Println()
}

func findAddends(ints []int, parts int, sum int) (product int, operation string, err error) {
	if parts < 0 || parts > 3 {
		return 0, "", errors.New("Number of parts not implemented")
	}

	l := len(ints)

	for i := 0; i < l; i++ {
		x := ints[i]
		if parts == 1 && x == sum {
			return x, fmt.Sprintf("%d", x), nil
		}

		for j := 0; j < l; j++ {
			y := ints[j]
			if parts == 2 && i != j && x+y == sum {
				return x * y, fmt.Sprintf("%d * %d", x, y), nil
			}

			for k := 0; k < l; k++ {
				z := ints[k]
				if parts == 3 && j != k && x+y+z == sum {
					return x * y * z, fmt.Sprintf("%d * %d * %d", x, y, z), nil
				}
			}
		}
	}

	return 0, "", errors.New("No matches found")
}
