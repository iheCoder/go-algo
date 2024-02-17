package search_a_2d_matrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	index := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	i, j := index/n, index%n
	if i >= m || j >= n {
		return false
	}
	return matrix[i][j] == target
}
