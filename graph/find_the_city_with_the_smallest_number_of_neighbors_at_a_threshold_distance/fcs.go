package find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	ans := []int{math.MaxInt32 / 2, -1}
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt32 / 2
		}
	}

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		mp[from][to], mp[to][from] = weight, weight
	}

	for k := 0; k < n; k++ {
		mp[k][k] = 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				mp[i][j] = minInt(mp[i][j], mp[i][k]+mp[k][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		var cnt int
		for j := 0; j < n; j++ {
			if mp[i][j] <= distanceThreshold {
				cnt++
			}
		}

		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}

	return ans[1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
