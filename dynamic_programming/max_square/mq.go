package max_square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	var ans int
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}
	}
	for j := 0; j < col; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ans = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] != '1' {
				continue
			}
			dp[i][j] = minInt(minInt(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			ans = maxInt(ans, dp[i][j])
		}
	}

	return ans * ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
