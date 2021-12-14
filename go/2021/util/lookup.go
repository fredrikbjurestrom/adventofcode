package util

import (
	"sort"
)

// Max returns the largest integer of inputs
func Max(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

// Min returns the largest integer of inputs
func Min(nums ...int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

// Mediam returns the median value
func Median(nums ...int) float64 {
	sort.Ints(nums)

	m := len(nums) / 2

	if m%2 > 0 {
		return float64(nums[m])
	}

	return float64((nums[m-1] + nums[m])) / 2
}

// Avg returns the average value
func Avg(nums ...int) float64 {
	s := 0
	l := len(nums)
	for _, v := range nums {
		s += v
	}

	return (float64(s)) / (float64(l))
}

// Contains check int slice if supplied value exists
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsString check int slice if supplied value exists
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DistinctStrings(s []string) []string {
	dict := map[string]int{}
	for _, a := range s {
		dict[a]++
	}

	result := make([]string, len(dict))
	i := 0
	for k := range dict {
		result[i] = k
		i++
	}

	return result
}
