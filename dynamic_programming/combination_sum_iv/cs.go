package combination_sum_iv

import "sort"

func combinationSum4Dep(nums []int, target int) int {
	sort.Ints(nums)
	edge := sort.SearchInts(nums, target)
	if edge < len(nums) && nums[edge] == target {
		edge++
	}
	nums = nums[:edge]

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		if m[i] {
			dp[i]++
		}
		for j := 1; j < i; j++ {
			if dp[j] <= 0 || dp[i-j] <= 0 {
				continue
			}
			dp[i] += dp[i-j]
		}
	}

	return dp[target]
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
