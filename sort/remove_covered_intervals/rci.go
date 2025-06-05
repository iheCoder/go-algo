package remove_covered_intervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0]) || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})

	last, ans := intervals[0], 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= last[0] && intervals[i][1] <= last[1] {
			continue
		}

		last = intervals[i]
		ans++
	}

	return ans
}
