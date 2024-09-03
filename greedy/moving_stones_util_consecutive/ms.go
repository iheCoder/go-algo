package moving_stones_util_consecutive

import "sort"

func numMovesStones(a int, b int, c int) []int {
	os := []int{a, b, c}
	sort.Ints(os)

	l, m, r := os[0], os[1], os[2]
	minMove := 2
	if m-l <= 2 || r-m <= 2 {
		minMove = 1
		if m-l == 1 && r-m == 1 {
			minMove = 0
		}
	}

	maxMove := r - l - 2

	return []int{minMove, maxMove}
}
