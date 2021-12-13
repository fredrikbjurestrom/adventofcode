package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day11 - Dumbo Octopus
func Day11() {
	fmt.Println("-------------------------")
	fmt.Println("DAY ELEVEN")

	input, err := util.FileAsStrings("./input/day11input.txt")
	if err != nil {
		panic(err)
	}

	grid := parseOctopusGrid(input)

	part1 := octopusFlashStep(grid, 100)

	part2grid := parseOctopusGrid(input)
	part2 := findSimultaneousFlashes(part2grid)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseOctopusGrid(input []string) [][]int {
	grid := make([][]int, len(input))

	for y, yv := range input {
		row := make([]int, len(yv))

		for x, xv := range strings.Split(yv, "") {
			n, _ := strconv.Atoi(xv)
			row[x] = n
		}

		grid[y] = row
	}

	return grid
}

func octopusFlashStep(grid [][]int, steps int) int {
	flashes := 0

	for i := 0; i < steps; i++ {
		for y, row := range grid {
			for x := range row {
				grid[y][x]++

				if grid[y][x] == 10 {
					flashOctopus(grid, y, x)
				}
			}
		}

		for y, row := range grid {
			for x, v := range row {
				if v > 9 {
					grid[y][x] = 0
					flashes++
				}
			}
		}
	}

	return flashes
}

func findSimultaneousFlashes(grid [][]int) int {
	i := 0
	for {
		i++

		for y, row := range grid {
			for x := range row {
				grid[y][x]++

				if grid[y][x] == 10 {
					flashOctopus(grid, y, x)
				}
			}
		}

		stepFlashes := 0
		for y, row := range grid {
			for x, v := range row {
				if v > 9 {
					grid[y][x] = 0
					stepFlashes++
				}
			}
		}

		if stepFlashes == len(grid)*len(grid[0]) {
			break
		}
	}

	return i
}

func flashOctopus(grid [][]int, y int, x int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) ||
				y+j < 0 ||
				y+j >= len(grid) ||
				x+i < 0 ||
				x+i >= len(grid[y]) {
				continue
			}

			grid[y+j][x+i]++

			if grid[y+j][x+i] == 10 {
				flashOctopus(grid, y+j, x+i)
			}
		}
	}
}
