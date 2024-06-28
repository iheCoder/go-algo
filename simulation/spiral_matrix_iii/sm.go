package spiral_matrix_iii

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	ans := [][]int{{rStart, cStart}}
	var dir int
	path, limit := 1, 2
	x, y := rStart, cStart
	for len(ans) < rows*cols {
		switch dir {
		case 0:
			if path >= limit {
				dir = 1
				path = 1
				continue
			}
			y++
			path++
			if x < rows && y < cols && x >= 0 && y >= 0 {
				ans = append(ans, []int{x, y})
			}
		case 1:
			if path >= limit {
				dir = 2
				path = 1
				limit++
				continue
			}
			x++
			path++
			if x < rows && y < cols && x >= 0 && y >= 0 {
				ans = append(ans, []int{x, y})
			}
		case 2:
			if path >= limit {
				dir = 3
				path = 1
				continue
			}
			y--
			path++
			if x < rows && y < cols && x >= 0 && y >= 0 {
				ans = append(ans, []int{x, y})
			}
		case 3:
			if path >= limit {
				dir = 0
				path = 1
				limit++
				continue
			}
			x--
			path++
			if x < rows && y < cols && x >= 0 && y >= 0 {
				ans = append(ans, []int{x, y})
			}
		}
	}

	return ans
}
