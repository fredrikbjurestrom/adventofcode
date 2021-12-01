package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2020/util"
)

// Day04 - Passport Processing
func Day04() {
	fmt.Println("-------------------------")
	fmt.Println("DAY FOUR")

	passports, err := readPassports("./input/day04input.txt")
	if err != nil {
		panic(err)
	}

	part1 := 0
	part2 := 0
	for _, passport := range passports {
		present, valid := validatePassport(passport)

		part1 += present
		part2 += valid
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func readPassports(path string) ([]map[string]string, error) {
	input, err := util.FileAsStrings(path)
	if err != nil {
		return nil, err
	}

	var result []map[string]string
	passport := make(map[string]string)
	for _, line := range input {
		if len(line) == 0 && passport != nil {
			result = append(result, passport)
			passport = make(map[string]string)
			continue
		}

		for _, v := range strings.Split(line, " ") {
			parts := strings.Split(v, ":")
			passport[parts[0]] = parts[1]
		}
	}

	result = append(result, passport)

	return result, nil
}

func validatePassport(passport map[string]string) (part1 int, part2 int) {
	present := 1
	valid := 1

	if byr, ok := passport["byr"]; ok {
		if !numberBetween(byr, 1920, 2002) {
			valid = 0
		}
	} else {
		present = 0
	}

	if iyr, ok := passport["iyr"]; ok {
		if !numberBetween(iyr, 2010, 2020) {
			valid = 0
		}
	} else {
		present = 0
	}

	if eyr, ok := passport["eyr"]; ok {
		if !numberBetween(eyr, 2020, 2030) {
			valid = 0
		}
	} else {
		present = 0
	}

	if hgt, ok := passport["hgt"]; ok {
		hgts := []rune(hgt)
		number := string(hgts[:len(hgts)-2])
		scale := string(hgts[len(hgts)-2:])

		if scale == "cm" {
			if !numberBetween(number, 150, 193) {
				valid = 0
			}
		} else if scale == "in" {
			if !numberBetween(number, 59, 76) {
				valid = 0
			}
		} else {
			valid = 0
		}
	} else {
		present = 0
	}

	if hcl, ok := passport["hcl"]; ok {
		if hcl != "amb" && hcl != "blu" && hcl != "brn" && hcl != "gry" && hcl != "grn" && hcl != "hzl" && hcl != "oth" {
			valid = 0
		}
	} else {
		present = 0
	}

	if _, ok := passport["ecl"]; ok {
		valid = valid
	} else {
		present = 0
	}
	if _, ok := passport["pid"]; ok {
		valid = valid
	} else {
		present = 0
	}

	return present, present * valid
}

func numberBetween(number string, from int64, to int64) bool {
	if n, err := strconv.ParseInt(number, 0, 16); err != nil {
		return n >= from && n <= to
	}

	return false
}
