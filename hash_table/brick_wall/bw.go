package brick_wall

func leastBricks(wall [][]int) int {
	n := len(wall)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		var sum int
		for j := 0; j < len(wall[i])-1; j++ {
			sum += wall[i][j]
			m[sum]++
		}
	}

	var crossedCount int
	for _, count := range m {
		if count > crossedCount {
			crossedCount = count
		}
	}
	return n - crossedCount
}
