package domino_and_tromino_tiling

func numTilings(n int) int {
	dp := make([][4]int, n+1)
	const mod int = 1e9 + 7
	dp[0][3] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][3]
		// 根据前面状态位置情况确定当前位置情况
		dp[i][1] = (dp[i-1][1] + dp[i-1][0]) % mod
		dp[i][2] = (dp[i-1][2] + dp[i-1][0]) % mod
		dp[i][3] = (((dp[i-1][1]+dp[i-1][2])%mod+dp[i-1][0])%mod + dp[i-1][3]) % mod
	}

	return dp[n][3]
}
