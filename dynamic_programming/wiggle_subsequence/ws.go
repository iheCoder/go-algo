package wiggle_subsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 1
	dp := make([][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i][0] = 1
		dp[i][1] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i][0] = maxInt(dp[i][0], dp[j][1]+1)
			} else if nums[i] < nums[j] {
				dp[i][1] = maxInt(dp[i][1], dp[j][0]+1)
			}
		}
		ans = maxInt(ans, maxInt(dp[i][0], dp[i][1]))
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
