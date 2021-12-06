package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (l Line) Distance() Point {
	return Point{l.end.x - l.start.x, l.end.y - l.start.y}
}

// Day05 - Hydrothermal Venture
func Day05() {
	fmt.Println("-------------------------")
	fmt.Println("DAY FIVE")

	input, err := util.FileAsStrings("./input/day05input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	horizontalAndDiagonalLines, err := parseLines(input, false)
	if err != nil {
		panic(err)
	}

	part1, err := findOverlaps(horizontalAndDiagonalLines)
	if err != nil {
		panic(err)
	}

	// part 2
	allLines, err := parseLines(input, true)
	if err != nil {
		panic(err)
	}

	part2, err := findOverlaps(allLines)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseLines(input []string, includeVertical bool) (lines []Line, err error) {
	for _, v := range input {
		parts := strings.Split(v, " -> ")
		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[1], ",")

		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		if includeVertical || (x1 == x2 || y1 == y2) {
			lines = append(lines, Line{Point{x1, y1}, Point{x2, y2}})
		}

	}

	return lines, nil
}

func findOverlaps(lines []Line) (result int, err error) {
	visited := make(map[Point]int)

	for _, line := range lines {
		d := line.Distance()
		xt := 0
		if d.x < 0 {
			xt = -1
		} else if d.x > 1 {
			xt = 1
		}

		yt := 0
		if d.y < 0 {
			yt = -1
		} else if d.y > 1 {
			yt = 1
		}

		p := line.start
		visited[p]++
		for {
			p.x += xt
			p.y += yt
			visited[p]++
			if p == line.end {
				break
			}
		}
	}

	for _, vis := range visited {
		if vis > 1 {
			result++
		}
	}

	return result, nil
}
