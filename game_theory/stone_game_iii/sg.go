package stone_game_iii

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	memo := make([]int, n)
	visited := make([]bool, n)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}

		if visited[i] {
			return memo[i]
		}

		var taken int
		best := math.MinInt
		for c := 1; c <= 3; c++ {
			j := i + c - 1
			if j >= n {
				break
			}

			taken += stoneValue[i+c-1]
			best = maxInt(best, taken-dfs(i+c))
		}

		memo[i] = best
		visited[i] = true
		return best
	}

	diff := dfs(0)
	switch {
	case diff > 0:
		return "Alice"
	case diff < 0:
		return "Bob"
	default:
		return "Tie"
	}
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
