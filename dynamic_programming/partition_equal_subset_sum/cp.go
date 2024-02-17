package partition_equal_subset_sum

import "sort"

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 > 0 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}

func canPartitionSet(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 > 0 {
		return false
	}
	target := sum / 2

	sort.Ints(nums)
	var f func(i, x int) bool
	f = func(i, x int) bool {
		if x > target {
			return false
		} else if x == target {
			return true
		}
		if i < 0 {
			return x == target
		}
		return f(i-1, x+nums[i]) || f(i-1, x)
	}
	return f(len(nums)-2, nums[len(nums)-1])
}
