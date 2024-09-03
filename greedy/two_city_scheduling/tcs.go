package two_city_scheduling

import "sort"

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	dp := make([]int, n)

	var ans int
	for i, cost := range costs {
		dp[i] = cost[1] - cost[0]
		ans += cost[0]
	}

	sort.Ints(dp)
	for i := 0; i < n/2; i++ {
		ans += dp[i]
	}

	return ans
}

// 	n := len(costs)
//	dp := make([]int, n)
//
//	idxA, idxB := make([]int, n), make([]int, n)
//	for i := 0; i < n; i++ {
//		idxA[i] = i
//		idxB[i] = i
//	}
//
//	sort.Slice(idxA, func(i, j int) bool {
//		return costs[i][0] < costs[j][0]
//	})
//
//	sort.Slice(idxB, func(i, j int) bool {
//		return costs[i][1] < costs[j][1]
//	})
//
//	var ans int
//	ca, cb := costs[idxA[0]][0], costs[idxB[0]][1]
//	if ca < cb || (ca == cb && costs[idxA[0]][1] > costs[idxB[0]][0]) {
//		dp[idxA[0]] = 1
//		ans += ca
//	} else {
//		dp[idxB[0]] = 1
//		ans += cb
//	}
