package edit_distance

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 若存在空串
	if m*n == 0 {
		return m + n
	}

	// 初始化dp
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			dp[i+1][j+1] = minInt(dp[i][j+1], dp[i+1][j]) + 1
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

//
//	// 计算dp值
//	for i := 1; i <= m; i++ {
//		for j := 1; j <= n; j++ {
//			sp := dp[i-1][j-1]
//			if word1[i-1] != word2[j-1] {
//				sp++
//			}
//			dp[i][j] = minInt(dp[i-1][j]+1, minInt(dp[i][j-1]+1, sp))
//		}
//	}
