package stone_game_ii

func stoneGameII260506(piles []int) int {
	n := len(piles)

	suffix := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + piles[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i >= n {
			return 0
		}

		if i+2*m >= n {
			return suffix[i]
		}

		if memo[i][m] != 0 {
			return memo[i][m]
		}

		best := 0
		for x := 1; x <= 2*m; x++ {
			nextM := maxInt(x, m)

			best = maxInt(best, suffix[i]-dfs(i+x, nextM))
		}

		memo[i][m] = best
		return best
	}

	return dfs(0, 1)
}
