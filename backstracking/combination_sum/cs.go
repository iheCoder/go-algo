package combination_sum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 1. 去除绝对不可能的
	sort.Ints(candidates)
	for len(candidates) > 0 && candidates[len(candidates)-1] > target {
		candidates = candidates[:len(candidates)-1]
	}

	// 2. 从0直到大于target进行回溯
	var cs func(i, t int, picks []int) [][]int
	cs = func(i, t int, picks []int) [][]int {
		if t == 0 {
			return [][]int{picks}
		} else if t < 0 {
			return nil
		}

		if i >= len(candidates) {
			return nil
		}

		var ans [][]int
		for c := 0; c*candidates[i] <= t; c++ {
			cp := make([]int, len(picks))
			copy(cp, picks)
			for j := 0; j < c; j++ {
				cp = append(cp, candidates[i])
			}
			a := cs(i+1, t-c*candidates[i], cp)
			ans = append(ans, a...)
		}
		return ans
	}

	return cs(0, target, []int{})
}
