package num_of_dice_rolls_with_target_sum

func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for x := 1; x <= target; x++ {
				if x-j < 0 {
					continue
				}

				dp[i][x] = (dp[i][x] + dp[i-1][x-j]) % mod
			}
		}
	}

	return dp[n][target]
}
