package longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == text2 {
		return len(text1)
	}

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}

			dp[i+1][j+1] = maxInt(dp[i][j+1], dp[i+1][j])
		}
	}

	return dp[m][n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
