package max_alternating_subsequence_sum

func maxAlternatingSum(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	even, odd := nums[0], 0
	for i := 1; i < len(nums); i++ {
		even = maxInt(even, odd+nums[i])
		odd = maxInt(odd, even-nums[i])
	}
	return int64(even)
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
