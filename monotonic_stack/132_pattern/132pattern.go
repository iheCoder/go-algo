package _32_pattern

import "math"

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	n := len(nums)
	ck := []int{nums[n-1]}
	mk := math.MinInt

	for i := n - 2; i >= 0; i-- {
		if nums[i] < mk {
			return true
		}

	}
}
