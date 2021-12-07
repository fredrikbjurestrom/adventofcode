package days

import (
	"fmt"
	"math"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day07 - The Treachery of Whales
func Day07() {
	fmt.Println("-------------------------")
	fmt.Println("DAY SEVEN")

	input, err := util.SeedAsInts("./input/day07input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	median := int(math.Round(util.Median(input...)))

	part1, err := calculateCrabFuelCost(input, median)
	if err != nil {
		panic(err)
	}

	// part 2
	rawAvg := util.Avg(input...)
	avgFloor := int(math.Floor(rawAvg))
	avgCeil := int(math.Floor(rawAvg))

	part2Floor, err := calculateIncrementalCrabFuelCost(input, avgFloor)
	if err != nil {
		panic(err)
	}

	part2Ceil, err := calculateIncrementalCrabFuelCost(input, avgCeil)
	if err != nil {
		panic(err)
	}

	part2 := part2Floor
	if part2Ceil < part2Floor {
		part2 = part2Ceil
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func findMedian(input []int) (result int, err error) {
	median := util.Median(input...)

	return int(math.Round(median)), nil
}

func calculateCrabFuelCost(input []int, target int) (result int, err error) {
	for _, v := range input {
		diff := int(math.Abs(float64(v - target)))
		result += diff
	}

	return result, nil
}

func calculateIncrementalCrabFuelCost(input []int, target int) (result int, err error) {
	for _, v := range input {
		diff := math.Abs(float64(v - target))

		cost := (diff / 2) * (1 + diff)

		result += int(math.Round(cost))
	}

	return result, nil
}
