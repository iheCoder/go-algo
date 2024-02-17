package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, left, right, bottom := -1, -1, n, m
	dir := 0
	var i, j int
	ans := []int{matrix[0][0]}
	for len(ans) < m*n {
		switch dir {
		case 0:
			j++
			if j >= right {
				top++
				j = right - 1
				dir = 1
				continue
			}

		case 1:
			i++
			if i >= bottom {
				right--
				i = bottom - 1
				dir = 2
				continue
			}

		case 2:
			j--
			if j <= left {
				bottom--
				j = left + 1
				dir = 3
				continue
			}

		case 3:
			i--
			if i <= top {
				left++
				i = top + 1
				dir = 0
				continue
			}

		}
		ans = append(ans, matrix[i][j])
	}
	return ans
}
