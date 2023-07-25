package unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < col && obstacleGrid[0][j] != 1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			left, top := 0, 0
			if j > 0 && obstacleGrid[i][j-1] != 1 {
				left = dp[i][j-1]
			}
			if i > 0 && obstacleGrid[i-1][j] != 1 {
				top = dp[i-1][j]
			}
			dp[i][j] = left + top
		}
	}

	return dp[row-1][col-1]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
