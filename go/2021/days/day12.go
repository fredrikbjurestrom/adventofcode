package days

import (
	"fmt"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

type Graph struct {
	edges map[string][]string
}

// Day12 - Passage Pathing
func Day12() {
	fmt.Println("-------------------------")
	fmt.Println("DAY TWELVE")

	input, err := util.FileAsStrings("./input/day12input.txt")
	if err != nil {
		panic(err)
	}

	g := parseGraph(input)

	part1 := findCavePaths(g, false)
	part2 := findCavePaths(g, true)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseGraph(input []string) Graph {
	g := Graph{}
	g.edges = map[string][]string{}

	for _, v := range input {
		parts := strings.Split(v, "-")
		n1 := parts[0]
		n2 := parts[1]

		g.edges[n1] = append(g.edges[n1], n2)
		g.edges[n2] = append(g.edges[n2], n1)
	}

	return g
}

func findCavePaths(g Graph, p2 bool) int {
	q := [][]string{{"start"}}
	paths := []string{}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		n1 := cur[len(cur)-1]

		if n1 == "end" {
			paths = append(paths, strings.Join(cur, ","))
			continue
		}

		// part 2 nonsense
		smallCaves := []string{}
		for _, c := range cur {
			if c == strings.ToLower(c) {
				smallCaves = append(smallCaves, c)
			}
		}

		d := false
		if len(smallCaves) > len(util.DistinctStrings(smallCaves)) {
			d = true
		}

		for _, n2 := range g.edges[n1] {
			if n2 == strings.ToUpper(n2) ||
				!util.ContainsString(cur, n2) ||
				(p2 && !d && n2 != "start") {

				next := make([]string, len(cur))
				copy(next, cur)

				next = append(next, n2)
				q = append(q, next)
			}
		}
	}

	return len(paths)
}
