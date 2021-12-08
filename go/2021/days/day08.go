package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

type Display struct {
	patterns [10]string
	output   [4]string
}

type SignalMap map[string]int

// Day08 - Seven Segment Search
func Day08() {
	fmt.Println("-------------------------")
	fmt.Println("DAY EIGHT")

	input, err := util.FileAsStrings("./input/day08input.txt")
	if err != nil {
		panic(err)
	}

	displays := parseDisplays(input)

	part1 := countEasyDigits(displays)
	part2 := sumOutput(displays)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func parseDisplays(input []string) []Display {
	output := make([]Display, len(input))

	for i, v := range input {
		parts := strings.Split(v, " | ")
		display := Display{}

		for pi, pattern := range strings.Fields(parts[0]) {
			display.patterns[pi] = pattern
		}

		for oi, output := range strings.Fields(parts[1]) {
			display.output[oi] = output
		}

		output[i] = display
	}

	return output
}

func countEasyDigits(displays []Display) int {
	result := 0

	for _, v := range displays {
		for _, o := range v.output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				result++
			}
		}
	}

	return result
}

func sumOutput(displays []Display) int {
	sum := 0
	for _, v := range displays {
		signalMap := getSignalMap(v.patterns)

		var outputStr string
		for _, o := range v.output {
			outputStr += strconv.Itoa(signalMap[SortString(o)])
		}

		output, _ := strconv.Atoi(outputStr)
		sum += output
	}

	return sum
}

func getSignalMap(pattern [10]string) SignalMap {
	var zero, one, two, three, four, five, six, seven, eight, nine string
	var twothreeorfive, zerosixornine []string

	for _, v := range pattern {
		switch len(v) {
		case 2:
			one = v
		case 3:
			seven = v
		case 4:
			four = v
		case 5:
			twothreeorfive = append(twothreeorfive, v)
		case 6:
			zerosixornine = append(zerosixornine, v)
		case 7:
			eight = v
		}
	}

	for _, v := range twothreeorfive {
		if strings.ContainsRune(v, rune(one[0])) && strings.ContainsRune(v, rune(one[1])) {
			three = v
		}
	}

	var b rune
	for _, v := range zerosixornine {
		if strings.ContainsRune(v, rune(three[0])) &&
			strings.ContainsRune(v, rune(three[1])) &&
			strings.ContainsRune(v, rune(three[2])) &&
			strings.ContainsRune(v, rune(three[3])) &&
			strings.ContainsRune(v, rune(three[4])) {
			nine = v

			for _, r := range nine {
				if !strings.ContainsRune(three, rune(r)) {
					b = r
				}
			}
		}
	}

	for _, v := range zerosixornine {
		if v == nine {
			continue
		}

		if strings.ContainsRune(v, rune(one[0])) && strings.ContainsRune(v, rune(one[1])) {
			zero = v
		} else {
			six = v
		}
	}

	for _, v := range twothreeorfive {
		if v == three {
			continue
		}

		if strings.ContainsRune(v, b) {
			five = v
		} else {
			two = v
		}
	}

	signalMap := SignalMap{
		SortString(zero):  0,
		SortString(one):   1,
		SortString(two):   2,
		SortString(three): 3,
		SortString(four):  4,
		SortString(five):  5,
		SortString(six):   6,
		SortString(seven): 7,
		SortString(eight): 8,
		SortString(nine):  9,
	}

	return signalMap
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
