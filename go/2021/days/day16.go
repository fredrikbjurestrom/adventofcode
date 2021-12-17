package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fredrikbjurestrom/adventofcode/go/2021/util"
)

// Day16 - Packet Decoder
func Day16() {
	fmt.Println("-------------------------")
	fmt.Println("DAY SIXTEEN")

	input, err := util.FileAsStrings("./input/day16input.txt")
	if err != nil {
		panic(err)
	}

	bits := hexToBits(input[0])
	part1, part2, _ := decodePacket(bits)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

	fmt.Println("-------------------------")
	fmt.Println()
}

func hexToBits(input string) []uint64 {
	bits := []uint64{}
	for _, hex := range strings.Split(input, "") {
		n, _ := strconv.ParseUint(hex, 16, 4)

		hexBits := []uint64{}
		for i := 0; i < 4; i++ {
			hexBits = append([]uint64{n & 0x1}, hexBits...)
			n = n >> 1
		}

		bits = append(bits, hexBits...)
	}

	return bits
}

func decodePacket(bits []uint64) (int, int64, int) {
	i := 0
	version := bitsToInt(bits[i : i+3])
	i += 3
	ptype := bitsToInt(bits[i : i+3])
	i += 3

	value := int64(0)

	if ptype == 4 {
		var litpos int
		value, litpos = literalValue(bits[i:])

		return int(version), value, i + litpos
	}

	lengthType := bits[i]
	i += 1

	versum := int(version)
	values := []int64{}
	if lengthType == 0 {
		bitcount := bitsToInt(bits[i : i+15])
		i += 15

		traversed := 0
		for traversed < int(bitcount) {
			subver, subval, subpos := decodePacket(bits[i+traversed:])

			versum += subver
			traversed += subpos
			values = append(values, subval)
		}
		i += int(bitcount)
	} else {
		pkgs := int(bitsToInt(bits[i : i+11]))
		i += 11

		for p := 0; p < pkgs; p++ {
			subver, subval, subpos := decodePacket(bits[i:])

			versum += subver
			i += subpos
			values = append(values, subval)
		}
	}

	switch ptype {
	case 0:
		value = util.Sum(values...)
	case 1:
		value = util.Product(values...)
	case 2:
		value = util.Min64(values...)
	case 3:
		value = util.Max64(values...)
	case 5:
		if values[0] > values[1] {
			value = 1
		}
	case 6:
		if values[0] < values[1] {
			value = 1
		}
	case 7:
		if values[0] == values[1] {
			value = 1
		}
	}

	return versum, value, i
}

func literalValue(bits []uint64) (int64, int) {
	agg := []uint64{}

	pos := 0
	for {
		grp := bits[pos : pos+5]

		agg = append(agg, grp[1:]...)
		pos += 5
		if grp[0] == 0 {
			break
		}
	}

	return bitsToInt(agg), pos
}

func bitsToInt(bits []uint64) int64 {
	bitstring := strings.Trim(strings.Join(strings.Split(fmt.Sprint(bits), " "), ""), "[]")
	i, _ := strconv.ParseInt(bitstring, 2, len(bitstring)+1)
	return i
}
