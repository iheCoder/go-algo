package num_of_longest_increasing_subsequence

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dc := make([]int, n)
	var maxCount int
	var ans int
	for i := 0; i < n; i++ {
		dp[i], dc[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}

			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				dc[i] = dc[j]
			} else if dp[j]+1 == dp[i] {
				dc[i] += dc[j]
			}
		}

		if maxCount == dp[i] {
			ans += dc[i]
		} else if dp[i] > maxCount {
			maxCount = dp[i]
			ans = dc[i]
		}
	}

	return ans
}
