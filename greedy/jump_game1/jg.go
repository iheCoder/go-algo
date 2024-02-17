package jump_game1

const maxSteps = 1<<31 - 1

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = maxSteps
	}
	dp[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		steps := nums[i]
		for j := 1; j <= steps && i+j < n; j++ {
			dp[i] = minInt(dp[i], dp[i+j]+1)
		}
	}
	return dp[0]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
