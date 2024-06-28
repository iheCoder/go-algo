package min_falling_path

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := matrix[0]
	var lastDP []int

	for i := 1; i < n; i++ {
		lastDP = dp
		dp = make([]int, n)
		for j := 0; j < n; j++ {
			dp[j] = lastDP[j] + matrix[i][j]
			if j > 0 {
				dp[j] = minInt(dp[j], lastDP[j-1]+matrix[i][j])
			}
			if j < n-1 {
				dp[j] = minInt(dp[j], lastDP[j+1]+matrix[i][j])
			}
		}
	}

	ans := dp[0]
	for i := 1; i < n; i++ {
		ans = minInt(ans, dp[i])
	}
	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
