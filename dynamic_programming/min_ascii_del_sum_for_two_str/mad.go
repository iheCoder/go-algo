package min_ascii_del_sum_for_two_str

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for i := 0; i <= n; i++ {
		if i > 0 {
			dp[0][i] = dp[0][i-1] + int(s2[i-1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = minInt(dp[i][j+1]+int(s1[i]), dp[i+1][j]+int(s2[j]))
			}
		}
	}

	return dp[m][n]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
