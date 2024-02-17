package longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
