package stone_game_ii

import "math"

func stoneGameII(piles []int) int {
	n := len(piles)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + piles[i-1]
	}

	type pair struct {
		x, m int
	}
	e := make(map[pair]int)
	var f func(i, m int) int
	f = func(i, m int) int {
		if i == n {
			return 0
		}

		p := pair{
			x: i,
			m: m,
		}
		if v, ok := e[p]; ok {
			return v
		}

		maxVal := math.MinInt
		for x := 1; x <= 2*m; x++ {
			if i+x > n {
				break
			}

			y := (prefixSum[i+x] - prefixSum[i]) - f(i+x, maxInt(x, m))
			maxVal = maxInt(maxVal, y)
		}

		e[p] = maxVal
		return maxVal
	}

	return (prefixSum[n] + f(0, 1)) / 2
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
