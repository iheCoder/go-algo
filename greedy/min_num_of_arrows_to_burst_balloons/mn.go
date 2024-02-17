package min_num_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1] || points[i][1] == points[j][1] && points[i][0] < points[j][0]
	})

	var p []int
	var ans, i int
	for i < len(points) {
		p = points[i]
		ans++
		i++
		for i < len(points) && points[i][0] <= p[1] {
			i++
		}
	}

	return ans
}
