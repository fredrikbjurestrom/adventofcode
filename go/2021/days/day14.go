package days

import (
	"fmt"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day14 - Extended Polymerization
func Day14() {
	fmt.Println("-------------------------")
	fmt.Println("DAY FOURTEEN")

	input, err := util.FileAsStrings("./input/day14input.txt")
	if err != nil {
		panic(err)
	}

	template, rules := parsePolymerization(input)

	part1 := naivePolymerScore(template, rules, 10)
	part2 := polymerScore(template, rules, 40)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parsePolymerization(input []string) (string, map[string]string) {
	template := input[0]
	rules := make(map[string]string, len(input)-2)

	for _, v := range input[2:] {
		parts := strings.Split(v, " -> ")
		rules[parts[0]] = parts[1]
	}

	return template, rules
}

func applyNaivePolymerization(template string, rules map[string]string) string {
	polymer := make([]string, (len(template)*2)-1)

	last := ""
	for i, v := range strings.Split(template, "") {
		pos := i * 2
		polymer[pos] = v

		if i > 0 {
			add := rules[last+v]
			polymer[pos-1] = add
		}

		last = v
	}

	return strings.Join(polymer, "")
}

func naivePolymerScore(template string, rules map[string]string, steps int) int {
	for i := 0; i < steps; i++ {
		template = applyNaivePolymerization(template, rules)
	}

	elements := map[string]int{}
	for _, v := range strings.Split(template, "") {
		elements[v]++
	}

	var min, max int
	for _, v := range elements {
		if min == 0 || v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func polymerScore(template string, rules map[string]string, steps int) int {
	pairs := map[string]int{}

	last := ""
	for i, v := range strings.Split(template, "") {

		if i > 0 {
			pairs[last+v]++
		}

		last = v
	}

	for i := 0; i < steps; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			keyParts := strings.Split(k, "")
			n := rules[k]
			newPairs[keyParts[0]+n] += v
			newPairs[n+keyParts[1]] += v
			delete(pairs, k)
		}

		for k, v := range newPairs {
			pairs[k] += v
		}
	}

	elements := map[string]int{}
	elements[template[0:1]] += 1

	for k, v := range pairs {
		keyParts := strings.Split(k, "")
		elements[keyParts[1]] += v
	}

	var min, max int
	for _, v := range elements {
		if min == 0 || v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}
