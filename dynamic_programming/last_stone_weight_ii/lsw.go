package last_stone_weight_ii

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, stone := range stones {
		for x := target; x >= stone; x-- {
			dp[x] = dp[x] || dp[x-stone]
		}
	}

	for i := target; i >= 0; i-- {
		if dp[i] {
			return sum - 2*i
		}
	}

	return 0
}
