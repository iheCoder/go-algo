package soup_serving

func soupServings(n int) float64 {
	n = (n + 24) / 25
	if n >= 179 {
		return 1
	}

	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
	}

	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}

		if a <= 0 {
			return 1
		}

		if b <= 0 {
			return 0
		}

		if dp[a][b] > 0 {
			return dp[a][b]
		}

		dp[a][b] = (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3)) / 4
		return dp[a][b]
	}

	return dfs(n, n)
}

func maxInt(a, b int) int {
	if b > a {
		return b
	}
	return a
}
