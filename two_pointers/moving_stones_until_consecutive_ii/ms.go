package moving_stones_until_consecutive_ii

import "sort"

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)

	n := len(stones)
	maxMove := maxInt(stones[n-1]-stones[1]-n+2, stones[n-2]-stones[0]-n+2)
	minMove := n
	var j int
	for i := 0; i < n; i++ {
		for j+1 < n && stones[j+1]-stones[i] <= n-1 {
			j++
		}

		if j-i+1 == n-1 && stones[j]-stones[i]+1 == n-1 {
			minMove = minInt(minMove, 2)
		} else {
			minMove = minInt(minMove, n-(j-i+1))
		}
	}

	return []int{minMove, maxMove}
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
