package reshape_the_matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := i*n + j
			ans[x/c][x%c] = mat[i][j]
		}
	}
	return ans
}
