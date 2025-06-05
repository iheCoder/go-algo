package count_square_submatrics_with_all_ones

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = minInt(minInt(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}

			ans += dp[i][j]
		}
	}

	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
