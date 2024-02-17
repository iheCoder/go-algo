package out_of_boundary

const mod int = 1e9 + 7

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	var ans int
	dp[0][startRow][startColumn] = 1
	for move := 0; move < maxMove; move++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				count := dp[move][i][j]
				if count <= 0 {
					continue
				}

				for _, dir := range dirs {
					ni, nj := i+dir.x, j+dir.y
					if ni >= 0 && nj >= 0 && ni < m && nj < n {
						dp[move+1][ni][nj] = (dp[move+1][ni][nj] + count) % mod
					} else {
						ans = (ans + count) % mod
					}
				}
			}
		}
	}

	return ans
}
