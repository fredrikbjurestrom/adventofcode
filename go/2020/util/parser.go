package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

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
