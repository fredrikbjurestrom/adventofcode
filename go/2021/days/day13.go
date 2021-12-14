package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day13 - Transparent Origami
func Day13() {
	fmt.Println("-------------------------")
	fmt.Println("DAY THIRTEEN")

	input, err := util.FileAsStrings("./input/day13input.txt")
	if err != nil {
		panic(err)
	}

	paper, instructions := parsePaper(input)

	part1instructions := []string{instructions[0]}
	firstFold := foldPaper(paper, part1instructions)

	fullFold := foldPaper(paper, instructions)

	fmt.Println("Part 1:", len(firstFold))
	fmt.Println("Part 2:")
	printPaper(fullFold)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parsePaper(input []string) ([]Point, []string) {
	paper := []Point{}
	instructions := []string{}

	for _, v := range input {
		if strings.HasPrefix(v, "fold") {
			instructions = append(instructions, v[11:])
		} else if len(v) > 0 {
			parts := strings.Split(v, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			paper = append(paper, Point{x, y})
		}
	}

	return paper, instructions
}

func foldPaper(paper []Point, instructions []string) []Point {
	foldedPaper := make([]Point, len(paper))
	copy(foldedPaper, paper)

	for _, instr := range instructions {
		instrParts := strings.Split(instr, "=")
		axis := instrParts[0]
		pos, _ := strconv.Atoi(instrParts[1])

		foldedMap := map[Point]bool{}

		for _, p := range foldedPaper {
			if (axis == "x" && p.x == pos) ||
				(axis == "y" && p.y == pos) {
				continue
			}

			if axis == "y" && p.y > pos {
				p.y -= (p.y - pos) * 2
			}

			if axis == "x" && p.x > pos {
				p.x -= (p.x - pos) * 2
			}

			foldedMap[p] = true
		}

		foldedPaper = make([]Point, len(foldedMap))
		i := 0
		for k := range foldedMap {
			foldedPaper[i] = k
			i++
		}
	}

	return foldedPaper
}

func printPaper(paper []Point) {
	mx := 0
	my := 0
	mp := map[Point]bool{}

	for _, p := range paper {
		if p.x > mx {
			mx = p.x
		}

		if p.y > my {
			my = p.y
		}

		mp[p] = true
	}

	for y := 0; y <= my; y++ {
		row := make([]string, mx+1)
		for x := 0; x <= mx; x++ {
			v := "."
			if (mp[Point{x, y}]) {
				v = "#"
			}
			row[x] = v
		}
		fmt.Println(strings.Join(row, ""))
	}
}
