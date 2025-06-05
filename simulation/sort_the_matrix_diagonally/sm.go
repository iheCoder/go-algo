package sort_the_matrix_diagonally

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	track := func(sx, sy int) {
		arr := make([]int, 0)
		x, y := sx, sy
		for {
			arr = append(arr, mat[x][y])
			x++
			y++
			if x >= m || y >= n {
				break
			}
		}
		sort.Ints(arr)
		x, y = sx, sy
		var i int
		for {
			mat[x][y] = arr[i]
			x++
			y++
			i++
			if x >= m || y >= n {
				break
			}
		}

	}

	for i := 0; i < m; i++ {
		track(i, 0)
	}
	for i := 1; i < n; i++ {
		track(0, i)
	}

	return mat
}
