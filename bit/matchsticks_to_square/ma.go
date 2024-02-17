package matchsticks_to_square

import "sort"

// 回溯法
func makesquare(matchsticks []int) bool {
	var sum int
	for _, matchstick := range matchsticks {
		sum += matchstick
	}
	if sum%4 > 0 {
		return false
	}
	aver := sum >> 2
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	edges := make([]int, 4)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx >= len(matchsticks) {
			return true
		}

		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= aver && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}

// 状态压缩+dp
func makesquareStateCompressDP(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}

	tLen := totalLen >> 2
	dp := make([]int, 1<<len(matchsticks))
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}

	for s := 1; s < len(dp); s++ {
		for k, v := range matchsticks {
			if s>>k&1 == 0 {
				continue
			}
			e := s &^ (1 << k)
			if dp[e] >= 0 && dp[e]+v <= tLen {
				dp[s] = (dp[e] + v) % tLen
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}
