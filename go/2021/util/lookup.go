package util

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

// Contens check int slice if supplied value exists
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
