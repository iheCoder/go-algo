package find_right_interval

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}

	sort.Slice(indexes, func(i, j int) bool {
		return intervals[indexes[i]][0] < intervals[indexes[j]][0]
	})

	ans := make([]int, 0, n)
	for _, interval := range intervals {
		r := sort.Search(n, func(i int) bool {
			return interval[1] <= intervals[indexes[i]][0]
		})
		if r >= n {
			r = -1
		} else {
			r = indexes[r]
		}
		ans = append(ans, r)
	}
	return ans
}
