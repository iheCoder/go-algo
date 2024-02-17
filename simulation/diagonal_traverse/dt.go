package diagonal_traverse

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) < 0 {
		return []int{}
	}

	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)

	var x, y int
	for i := 0; i < m+n-1; i++ {
		if i%2 == 0 {
			x = min(m-1, i)
			// 之所以y是i-m+1的原因是找规律吧！
			y = max(0, i-m+1)
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		} else {
			x = max(i-n+1, 0)
			y = min(n-1, i)
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
