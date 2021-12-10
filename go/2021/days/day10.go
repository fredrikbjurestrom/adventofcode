package days

import (
	"fmt"
	"math"
	"sort"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day10 - Syntax Scoring
func Day10() {
	fmt.Println("-------------------------")
	fmt.Println("DAY TEN")

	input, err := util.FileAsStrings("./input/day10input.txt")
	if err != nil {
		panic(err)
	}

	part1, part2 := checkLineSyntax(input)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func checkLineSyntax(input []string) (errorScore int, completionScore int) {
	completionScores := []int{}

	for _, v := range input {
		valid, char, completion := syntaxValidator(v)
		rowCompletion := 0

		if !valid {
			switch char {
			case ')':
				errorScore += 3
			case ']':
				errorScore += 57
			case '}':
				errorScore += 1197
			case '>':
				errorScore += 25137
			}
		} else {
			for _, c := range completion {
				rowCompletion *= 5
				switch rune(c) {
				case ')':
					rowCompletion += 1
				case ']':
					rowCompletion += 2
				case '}':
					rowCompletion += 3
				case '>':
					rowCompletion += 4
				}
			}

			completionScores = append(completionScores, rowCompletion)
		}

	}

	sort.Ints(completionScores)
	middle := int(math.Floor(float64(len(completionScores)) / 2))

	return errorScore, completionScores[middle]
}

func syntaxValidator(line string) (bool, rune, string) {
	tree := map[int]rune{}
	depth := -1
	for _, v := range line {
		if rune(v) == '(' || rune(v) == '[' || rune(v) == '{' || rune(v) == '<' {
			depth++
			tree[depth] = rune(v)
			continue
		}

		if (rune(v) == ')' && tree[depth] != '(') ||
			(rune(v) == ']' && tree[depth] != '[') ||
			(rune(v) == '}' && tree[depth] != '{') ||
			(rune(v) == '>' && tree[depth] != '<') {
			return false, rune(v), ""
		}

		delete(tree, depth)
		depth--
	}

	completion := ""
	for d := depth; d > -1; d-- {
		switch tree[d] {
		case '(':
			completion += ")"
		case '[':
			completion += "]"
		case '{':
			completion += "}"
		case '<':
			completion += ">"
		}
	}

	return true, -1, completion
}
