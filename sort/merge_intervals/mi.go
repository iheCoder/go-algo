package merge_intervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1]
	})

	left := intervals[0][0]
	right := intervals[0][1]
	n := len(intervals)
	var ans [][]int
	for i := 1; i < n; i++ {
		if intervals[i][0] <= right {
			if intervals[i][1] > right {
				right = intervals[i][1]
			}
			continue
		}
		ans = append(ans, []int{left, right})
		left, right = intervals[i][0], intervals[i][1]
	}
	ans = append(ans, []int{left, right})

	return ans
}

// 无法应对[1,2],[3,4],[1,10]
// 	left := intervals[0][0]
//	rightEdge := intervals[0][1]
//	n := len(intervals)
//	var ans [][]int
//	for i := 1; i < n; i++ {
//		if left > intervals[i][0] {
//			left = intervals[i][0]
//		}
//		if intervals[i][0] <= rightEdge {
//			rightEdge = intervals[i][1]
//			continue
//		}
//		ans = append(ans, []int{left, intervals[i-1][1]})
//		left = intervals[i][0]
//	}
//	ans = append(ans, []int{left, intervals[n-1][1]})
//
//	return ans
