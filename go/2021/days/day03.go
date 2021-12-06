package days

import (
	"fmt"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day03 - Binary Diagnostic
func Day03() {
	fmt.Println("-------------------------")
	fmt.Println("DAY THREE")

	input, bitSize, err := util.BinaryFileAsInts("./input/day03input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	powerConsumption, err := getPowerConsumption(input, bitSize)
	if err != nil {
		panic(err)
	}

	// part 2
	oxygenGeneratorRating, err := getOxygenGeneratorRating(input, bitSize)
	if err != nil {
		panic(err)
	}

	co2ScrubberRating, err := getCo2ScrubberRating(input, bitSize)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", powerConsumption)
	fmt.Println("Part 2:", oxygenGeneratorRating*co2ScrubberRating)

	fmt.Println("-------------------------")
	fmt.Println()
}

func getPowerConsumption(input []uint64, bitSize uint) (result uint64, err error) {
	naiveCounter := make([]int, bitSize)

	var epsilon uint64 = 0
	var gamma uint64 = 0

	for _, n := range input {
		for i, _ := range naiveCounter {
			if n&(1<<i) > 0 {
				naiveCounter[i] += 1
			}
		}
	}

	for i, n := range naiveCounter {
		if n >= len(input)/2 {
			epsilon |= (1 << i)
		} else {
			gamma |= (1 << i)
		}
	}

	return epsilon * gamma, nil
}

func getOxygenGeneratorRating(input []uint64, bitSize uint) (result uint64, err error) {
	for pos := bitSize - 1; pos >= 0; pos-- {
		input = filterRatings(input, true, pos)

		if len(input) < 2 {
			result = input[0]
			break
		}
	}

	return result, nil
}

func getCo2ScrubberRating(input []uint64, bitSize uint) (result uint64, err error) {
	for pos := bitSize - 1; pos >= 0; pos-- {
		input = filterRatings(input, false, pos)

		if len(input) < 2 {
			result = input[0]
			break
		}
	}

	return result, nil
}

func filterRatings(input []uint64, most bool, position uint) (result []uint64) {
	set := 0
	unset := 0
	for _, n := range input {
		if n&(1<<position) > 0 {
			set += 1
		} else {
			unset += 1
		}
	}

	for _, n := range input {
		if set >= unset {
			if most && n&(1<<position) > 0 {
				result = append(result, n)
			}

			if !most && n&(1<<position) == 0 {
				result = append(result, n)
			}
		} else {
			if most && n&(1<<position) == 0 {
				result = append(result, n)
			}

			if !most && n&(1<<position) > 0 {
				result = append(result, n)
			}
		}
	}

	return result
}
