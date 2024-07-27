package min_cost_for_tickets

import "math"

func mincostTickets(days []int, costs []int) int {
	memo := [366]int{}
	durations := []int{1, 7, 30}

	var dp func(idx int) int
	dp = func(idx int) int {
		if idx >= len(days) {
			return 0
		}

		if memo[idx] > 0 {
			return memo[idx]
		}

		memo[idx] = math.MaxInt32
		j := idx
		for i := 0; i < 3; i++ {
			// dp(j) 是从下一个不在当前票有效期内的第 j 天开始的最小成本
			for ; j < len(days) && days[j] < days[idx]+durations[i]; j++ {
			}
			memo[idx] = minInt(memo[idx], costs[i]+dp(j))
		}
		return memo[idx]
	}

	return dp(0)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
