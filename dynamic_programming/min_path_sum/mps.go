package min_path_sum

const maxPath = 1 << 31

func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			left, top := maxPath, maxPath
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 {
				top = dp[i-1][j]
			}
			dp[i][j] = minInt(left, top) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func minInt(x, y int) int {
	if x == maxPath && y == maxPath {
		return 0
	}
	if x == maxPath {
		return y
	}
	if y == maxPath {
		return x
	}
	if x < y {
		return x
	}
	return y
}
