package life_game

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	cb := make([][]int, m)
	for i := 0; i < m; i++ {
		cb[i] = make([]int, n)
		copy(cb[i], board[i])
	}

	var cf, gc func(i, j int) int
	gc = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if i >= m || j >= n {
			return 0
		}
		return cb[i][j]
	}
	cf = func(i, j int) int {
		var c int
		c += gc(i-1, j-1) + gc(i-1, j) + gc(i-1, j+1)
		c += gc(i, j-1) + gc(i, j+1)
		c += gc(i+1, j-1) + gc(i+1, j) + gc(i+1, j+1)
		return c
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := cf(i, j)
			if c < 2 || c > 3 {
				board[i][j] = 0
			} else if c == 3 {
				board[i][j] = 1
			}
		}
	}
}
