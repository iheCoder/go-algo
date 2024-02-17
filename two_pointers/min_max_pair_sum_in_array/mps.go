package min_max_pair_sum_in_array

import "sort"

func minPairSum(nums []int) int {
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}

	sort.Ints(nums)

	var r int
	i, j := 0, len(nums)-1
	for i < j {
		r = maxInt(r, nums[i]+nums[j])
		i++
		j--
	}

	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
