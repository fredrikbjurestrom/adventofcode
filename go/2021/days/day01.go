package days

import (
	"fmt"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day01 - Sonar Sweep
func Day01() {
	fmt.Println("-------------------------")
	fmt.Println("DAY ONE")

	ints, err := util.FileAsInts("./input/day01input.txt")
	if err != nil {
		panic(err)
	}

	part1, err := countIncrements(ints)
	if err != nil {
		panic(err)
	}

	part2, err := countSlidingWindowIncrements(ints)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func countIncrements(ints []int) (increments int, err error) {
	l := len(ints)

	if l < 2 {
		return 0, nil
	}

	for i := 1; i < l; i++ {
		if ints[i] > ints[i-1] {
			increments += 1
		}
	}

	return increments, nil
}

func countSlidingWindowIncrements(ints []int) (increments int, err error) {
	l := len(ints)

	if l < 4 {
		return 0, nil
	}

	for i := 3; i < l; i++ {
		curWindow := ints[i-2 : i+1]
		prevWindow := ints[i-3 : i]

		if sum(curWindow) > sum(prevWindow) {
			increments += 1
		}
	}

	return increments, nil
}

func sum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}
