package days

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

type ValuedGraph struct {
	edges map[Point]map[Point]int
}

type Item struct {
	value    Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Day15 - Chiton
func Day15() {
	fmt.Println("-------------------------")
	fmt.Println("DAY FIFTEEN")

	input, err := util.FileAsStrings("./input/day15input.txt")
	if err != nil {
		panic(err)
	}

	g1, g1end := parseValuedGraph(input, 1)
	part1 := shortestPath(g1, g1end)

	g2, g2end := parseValuedGraph(input, 5)
	part2 := shortestPath(g2, g2end)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseValuedGraph(input []string, multiplier int) (ValuedGraph, Point) {
	g := ValuedGraph{}
	g.edges = map[Point]map[Point]int{}

	my := len(input)
	mx := len(input[0])

	end := Point{mx*multiplier - 1, my*multiplier - 1}

	for y := 0; y < my*multiplier; y++ {
		row := strings.Split(input[y%my], "")
		for x := 0; x < mx*multiplier; x++ {
			n1 := Point{x, y}
			v := row[x%mx]

			baseWeight, _ := strconv.Atoi(v)
			weight := (baseWeight + (y / my) + (x / mx))

			if weight > 9 {
				weight = weight % 9
			}

			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if (i == 0 && j == 0) ||
						y+j < 0 ||
						y+j >= my*multiplier ||
						x+i < 0 ||
						x+i >= mx*multiplier ||
						(i != 0 && j != 0) {
						continue
					}

					n2 := Point{x + i, y + j}

					if g.edges[n2] == nil {
						g.edges[n2] = map[Point]int{}
					}

					g.edges[n2][n1] = weight
				}
			}
		}
	}

	return g, end
}

func shortestPath(g ValuedGraph, end Point) int {
	start := Point{0, 0}

	visited := map[Point]bool{}

	score := map[Point]int{}
	for k := range g.edges {
		score[k] = math.MaxInt64
	}

	score[start] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    Point{0, 0},
		priority: 0,
		index:    0,
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		n1 := heap.Pop(&pq).(*Item)

		visited[n1.value] = true

		for n2, weight := range g.edges[n1.value] {
			if visited[n2] {
				continue
			}

			newScore := score[n1.value] + weight
			oldScore := score[n2]

			if newScore < oldScore {
				score[n2] = newScore

				next := &Item{
					value:    n2,
					priority: -newScore,
				}

				heap.Push(&pq, next)
			}
		}
	}

	return score[end]
}
