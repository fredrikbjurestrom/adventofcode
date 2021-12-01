package days

import (
	"fmt"

	"github.com/fredrikbjurestrom/adventofcode/go/2020/util"
)

// Day03 - Toboggan Trajectory
func Day03() {
	fmt.Println("-------------------------")
	fmt.Println("DAY THREE")

	input, err := util.FileAsStrings("./input/day03input.txt")
	if err != nil {
		panic(err)
	}

	rightOneDownOne := 0
	rightThreeDownOne := 0
	rightFiveDownOne := 0
	rightSevenDownOne := 0
	rightOneDownTwo := 0

	for y, v := range input {

		rightOneDownOne += tree(v, y*1)
		rightThreeDownOne += tree(v, y*3)
		rightFiveDownOne += tree(v, y*5)
		rightSevenDownOne += tree(v, y*7)

		if y%2 == 0 {
			rightOneDownTwo += tree(v, (y*1)/2)
		}
	}

	part2 := rightOneDownOne * rightThreeDownOne * rightFiveDownOne * rightSevenDownOne * rightOneDownTwo

	fmt.Println("Part 1:", rightThreeDownOne)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func tree(row string, x int) int {
	pos := x % len(row)

	if string(row[pos]) == "#" {
		return 1
	}

	return 0
}
