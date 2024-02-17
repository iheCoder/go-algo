package triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	var x, last int
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			x = triangle[i][j]
			if j > 0 && j < i {
				x += minInt(dp[j], last)
			} else if j > 0 {
				x += last
			} else {
				x += dp[j]
			}
			last = dp[j]
			dp[j] = x
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
