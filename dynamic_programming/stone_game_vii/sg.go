package stone_game_vii

func stoneGameVII(stones []int) int {
	n := len(stones)

	pre := make([]int, n+1)
	for i, stone := range stones {
		pre[i+1] = pre[i] + stone
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for l := 0; l <= n-length; l++ {
			r := l + length - 1

			scoreL := pre[r+1] - pre[l+1]
			scoreR := pre[r] - pre[l]
			dp[l][r] = maxInt(
				scoreL-dp[l+1][r],
				scoreR-dp[l][r-1],
			)
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
