package path_with_max_gold

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var ans int
	var f func(i, j, cur int)
	f = func(i, j, cur int) {
		val := grid[i][j]
		grid[i][j] = 0

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= n || grid[ni][nj] <= 0 {
				continue
			}

			f(ni, nj, cur+grid[ni][nj])
		}

		grid[i][j] = val
		ans = maxInt(ans, cur)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				f(i, j, grid[i][j])
			}
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
