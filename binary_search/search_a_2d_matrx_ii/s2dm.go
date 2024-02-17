package search_a_2d_matrx_ii

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	lx, hx := 0, m-1
	ly, hy := 0, n-1

	for lx <= hx || ly <= hy {
		mx, my := lx+(hx-lx)/2, ly+(hy-ly)/2
		if matrix[mx][my] > target {
			hx, hy = mx-1, my-1
		} else if matrix[mx][my] < target {
			lx, ly = mx+1, my+1
		} else {
			return true
		}
	}
	return false
}

func searchMatrixZ(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
