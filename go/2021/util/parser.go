package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// BinaryFileAsInts reads a file at supplied location and returns an uint64 array. If there's an error, it
// returns the uints successfully read so far as well along with identified bitSize and the error value.
func BinaryFileAsInts(path string) ([]uint64, uint, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	return readBinary(f)
}

// FileAsInts reads a file at supplied location and returns an int array. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func FileAsInts(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readInts(f)
}

// SeedAsInts reads a comma separated seed at supplied location and returns an int array. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func SeedAsInts(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readSeedInts(f)
}

// FileAsStrings reads a file at supplied location and returns a string array. If there's an error, it
// returns the strings successfully read so far as well as the error value.
func FileAsStrings(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return readStrings(f)
}

func readBinary(r io.Reader) ([]uint64, uint, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []uint64
	var bitSize int
	for scanner.Scan() {
		row := scanner.Text()

		if len(row) > bitSize {
			bitSize = len(row)
		}

		x, err := strconv.ParseUint(scanner.Text(), 2, 64)
		if err != nil {
			return result, uint(bitSize), err
		}
		result = append(result, x)
	}
	return result, uint(bitSize), scanner.Err()
}

func readStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func readSeedInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		seed := scanner.Text()
		for _, strv := range strings.Split(seed, ",") {
			num, err := strconv.Atoi(strv)
			if err != nil {
				return result, err
			}
			result = append(result, num)
		}
	}
	return result, scanner.Err()
}
