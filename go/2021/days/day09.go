package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

type Location struct {
	pos   Point
	value int
}

// Day09 - Smoke Basin
func Day09() {
	fmt.Println("-------------------------")
	fmt.Println("DAY NINE")

	input, err := util.FileAsStrings("./input/day09input.txt")
	if err != nil {
		panic(err)
	}

	heightMap := parseHeightMap(input)
	lowPoints := getLowPoints(heightMap)

	part1 := getRiskLevel(lowPoints)
	part2 := getBasinScore(heightMap, lowPoints)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseHeightMap(input []string) [][]int {
	heightMap := make([][]int, len(input))

	for y, row := range input {
		rowMap := make([]int, len(row))

		for x, v := range strings.Split(row, "") {
			iv, _ := strconv.Atoi(v)
			rowMap[x] = iv
		}
		heightMap[y] = rowMap
	}

	return heightMap
}

func getLowPoints(heightMap [][]int) []Location {
	var lowPoints []Location

	for y, row := range heightMap {
		ylow := y == 0
		yhigh := y == len(heightMap)-1

		for x, v := range row {
			xlow := x == 0
			xhigh := x == len(row)-1

			if (ylow || v < heightMap[y-1][x]) &&
				(yhigh || v < heightMap[y+1][x]) &&
				(xlow || v < heightMap[y][x-1]) &&
				(xhigh || v < heightMap[y][x+1]) {
				lowPoints = append(lowPoints, Location{Point{x, y}, v})
			}
		}
	}

	return lowPoints
}

func getRiskLevel(lowPoints []Location) int {
	s := 0
	for _, v := range lowPoints {
		s += v.value + 1
	}

	return s
}

func getBasinScore(heightMap [][]int, lowPoints []Location) int {
	basinSizes := []int{}

	for _, v := range lowPoints {
		if basinPoints := traverseBasinPoint(heightMap, v.pos, map[Point]bool{}); len(basinPoints) > 0 {
			basinSizes = append(basinSizes, len(basinPoints))
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func traverseBasinPoint(heightMap [][]int, startPos Point, visited map[Point]bool) []Point {
	basinPoints := []Point{}

	if visited[startPos] {
		return []Point{}
	}

	visited[startPos] = true

	for up := startPos.y + 1; up < len(heightMap); up++ {
		if heightMap[up][startPos.x] == 9 {
			break
		}

		basinPoints = append(basinPoints, Point{startPos.x, up})
	}

	for down := startPos.y - 1; down > -1; down-- {
		if heightMap[down][startPos.x] == 9 {
			break
		}

		basinPoints = append(basinPoints, Point{startPos.x, down})
	}

	for right := startPos.x + 1; right < len(heightMap[startPos.y]); right++ {
		if heightMap[startPos.y][right] == 9 {
			break
		}

		basinPoints = append(basinPoints, Point{right, startPos.y})
	}

	for left := startPos.x - 1; left > -1; left-- {
		if heightMap[startPos.y][left] == 9 {
			break
		}

		basinPoints = append(basinPoints, Point{left, startPos.y})
	}

	for _, v := range basinPoints {
		basinPoints = append(basinPoints, traverseBasinPoint(heightMap, v, visited)...)
	}

	basinPointMap := map[Point]bool{}
	uniqueBasinPoints := []Point{}
	for _, v := range basinPoints {
		if _, match := basinPointMap[v]; !match {
			basinPointMap[v] = true
			uniqueBasinPoints = append(uniqueBasinPoints, v)
		}
	}

	return uniqueBasinPoints
}
