package score_after_flipping_matrix

// 之所以后面的不能够通过相加大于前面的在于，如果可以那么也可以进行翻转列
func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 1 << (n - 1) * m
	for i := 1; i < n; i++ {
		var ones int
		for _, row := range grid {
			if row[i] == row[0] {
				ones++
			}
		}

		if ones < m-ones {
			ones = m - ones
		}

		ans += 1 << (n - 1 - i) * ones
	}

	return ans
}
