package largest_divisible_subset

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	maxSize, maxVal := 1, nums[0]
	for i, num := range nums {
		for j := 0; j < i; j++ {
			if num%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > maxSize {
					maxSize = dp[i]
					maxVal = num
				}
			}
		}
	}

	res := make([]int, 0, maxSize)
	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxSize--
			maxVal = nums[i]
		}
	}

	return res
}
