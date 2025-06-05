package get_equal_substr_with_budget

import "sort"

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	diffs := make([]int, n+1)
	for i := 0; i < n; i++ {
		diffs[i+1] = diffs[i] + diff(s[i], t[i])
	}

	maxLen := 0
	for i := 1; i <= n; i++ {
		target := diffs[i] - maxCost
		start := sort.Search(n, func(i int) bool {
			return diffs[i] >= target
		})
		maxLen = maxInt(maxLen, i-start)
	}

	return maxLen
}

func diff(a, b byte) int {
	d := int(a) - int(b)
	if d < 0 {
		d = -d
	}

	return d
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
