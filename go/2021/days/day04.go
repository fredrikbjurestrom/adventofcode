package days

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day04 - Giant Squid
func Day04() {
	fmt.Println("-------------------------")
	fmt.Println("DAY FOUR")

	input, err := util.FileAsStrings("./input/day04input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	draw, boards, boardSum, err := parseBoards(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(boardSum)

	bestScore, err := scoreBoards(draw, boards, boardSum, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(boardSum)

	// part 2
	worstScore, err := scoreBoards(draw, boards, boardSum, false)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", bestScore)
	fmt.Println("Part 2:", worstScore)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseBoards(input []string) (draw []int, boards [][5][5]int, boardSum []int, err error) {
	drawStr := strings.Split(input[0], ",")
	drawInt := make([]int, len(drawStr))
	for i, v := range drawStr {
		drawInt[i], _ = strconv.Atoi(v)
	}

	boardCount := len(input) / 6
	boardNo := 0
	parsedBoards := make([][5][5]int, boardCount)
	parsedBoardSum := make([]int, boardCount)
	for i := 2; i < len(input); i++ {
		if len(input[i]) == 0 {
			boardNo++
			continue
		}

		y := (i - 2) - boardNo*6
		for x, v := range strings.Fields(input[i]) {
			value, _ := strconv.Atoi(v)
			parsedBoards[boardNo][y][x] = value
			parsedBoardSum[boardNo] += value
		}
	}

	return drawInt, parsedBoards, parsedBoardSum, nil
}

func scoreBoards(draw []int, boards [][5][5]int, boardSum []int, pickBest bool) (result int, err error) {
	xmarks := make([][5]int, len(boards))
	ymarks := make([][5]int, len(boards))
	unmarkedSum := make([]int, len(boardSum))
	copy(unmarkedSum, boardSum)
	boardScore := make([]int, len(boards))

	matchedBoards := 0
	for _, no := range draw {
		for boardId, board := range boards {
			for y, row := range board {
				for x, v := range row {
					if v == no {
						if boardScore[boardId] > 0 {
							continue
						}

						unmarkedSum[boardId] -= no

						xmarks[boardId][x]++
						ymarks[boardId][y]++

						if xmarks[boardId][x] == 5 || ymarks[boardId][y] == 5 {
							matchedBoards++
							boardScore[boardId] = unmarkedSum[boardId] * no

							if pickBest {
								return boardScore[boardId], nil
							}

							if !pickBest && matchedBoards == len(boards) {
								return boardScore[boardId], nil
							}
						}
					}
				}
			}
		}
	}

	return 0, errors.New("No winning board found")
}
