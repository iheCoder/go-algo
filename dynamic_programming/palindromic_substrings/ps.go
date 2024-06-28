package palindromic_substrings

type pair struct {
	count            int
	isNotPalindromic bool
}

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]pair, n)
		dp[i][i] = pair{
			count:            1,
			isNotPalindromic: false,
		}
	}

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			dp[i][j] = pair{
				count:            dp[i+1][j].count + dp[i][j-1].count - dp[i+1][j-1].count,
				isNotPalindromic: true,
			}
			if s[i] == s[j] && !dp[i+1][j-1].isNotPalindromic {
				dp[i][j].count++
				dp[i][j].isNotPalindromic = false
				continue
			}
		}
	}
	return dp[0][n-1].count
}
