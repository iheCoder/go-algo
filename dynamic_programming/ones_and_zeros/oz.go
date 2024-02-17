package ones_and_zeros

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	dp := make([][][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 0; i < l; i++ {
		zeros := strings.Count(strs[i], "0")
		ones := len(strs[i]) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				// 若当前数量满足当前容量
				if j >= zeros && k >= ones {
					dp[i+1][j][k] = maxInt(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}

	return dp[l][m][n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
