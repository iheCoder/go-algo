package matrix_block_sum

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	matrixSum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrixSum[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			matrixSum[i][j] = matrixSum[i-1][j] + matrixSum[i][j-1] - matrixSum[i-1][j-1] + mat[i-1][j-1]
		}
	}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	get := func(x, y int) int {
		if x < 0 {
			x = 0
		}
		if x > m {
			x = m
		}
		if y < 0 {
			y = 0
		}
		if y > n {
			y = n
		}

		return matrixSum[x][y]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = get(i+k+1, j+k+1) -
				get(i-k, j+k+1) -
				get(i+k+1, j-k) +
				get(i-k, j-k)
		}
	}

	return ans
}
