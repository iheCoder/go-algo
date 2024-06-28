package partition_array_into_disjoint_intervals

import "math"

// 从左到右的最大值，从右到左的最小值
func partitionDisjoint(nums []int) int {
	n := len(nums)
	lm, rm := make([]int, n), make([]int, n)
	mx := -1
	for i := 0; i < n; i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
		lm[i] = mx
	}

	mx = math.MaxInt
	for i := n - 1; i >= 0; i-- {
		if nums[i] < mx {
			mx = nums[i]
		}
		rm[i] = mx
	}

	for i := 1; i < n; i++ {
		if lm[i-1] <= rm[i] {
			return i
		}
	}
	return 0
}
