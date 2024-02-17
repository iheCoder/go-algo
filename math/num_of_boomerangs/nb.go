package num_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}

	n := len(points)
	var ans int
	for i := 0; i < n; i++ {
		cnts := make(map[int]int)
		for j := 0; j < n; j++ {
			cnts[calDistance(points[i], points[j])]++
		}

		for _, cnt := range cnts {
			ans += cnt * (cnt - 1)
		}
	}

	return ans
}

func calDistance(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
