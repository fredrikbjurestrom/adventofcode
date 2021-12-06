package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day06 - Lanternfish
func Day06() {
	fmt.Println("-------------------------")
	fmt.Println("DAY SIX")

	input, err := util.FileAsStrings("./input/day06input.txt")
	if err != nil {
		panic(err)
	}

	population, err := initialFishPopulation(input)
	if err != nil {
		panic(err)
	}

	part1, err := simulatePopulationGrowth(population, 80)
	if err != nil {
		panic(err)
	}

	part2, err := simulatePopulationGrowth(population, 256)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func initialFishPopulation(input []string) (result [9]int, err error) {
	pop := [9]int{}

	for _, line := range input {
		for _, strv := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(strv)
			pop[num]++
		}
	}

	return pop, nil
}

func simulatePopulationGrowth(population [9]int, days int) (result int, err error) {
	for d := days; d > 0; d-- {
		parents := population[0]

		population[0] = population[1]
		population[1] = population[2]
		population[2] = population[3]
		population[3] = population[4]
		population[4] = population[5]
		population[5] = population[6]
		population[6] = population[7] + parents
		population[7] = population[8]
		population[8] = parents
	}

	for _, v := range population {
		result += v
	}

	return result, nil
}
