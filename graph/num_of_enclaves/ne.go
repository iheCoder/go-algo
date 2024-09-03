package num_of_enclaves

func numEnclaves(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0

		for _, d := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			dfs(i+d[0], j+d[1])
		}

		return 1
	}

	for i := 0; i < len(grid); i++ {
		dfs(i, 0)
		dfs(i, len(grid[0])-1)
	}

	for j := 0; j < len(grid[0]); j++ {
		dfs(0, j)
		dfs(len(grid)-1, j)
	}

	var res int
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				res++
			}
		}
	}

	return res
}
