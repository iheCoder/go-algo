package surrounded_regions

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	pos := make([][]int, 0)
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			pos = append(pos, []int{0, i})
		}
		if board[m-1][i] == 'O' {
			pos = append(pos, []int{m - 1, i})
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			pos = append(pos, []int{i, 0})
		}
		if board[i][n-1] == 'O' {
			pos = append(pos, []int{i, n - 1})
		}
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		board[x][y] = 'Y'
		if x > 0 && board[x-1][y] == 'O' {
			dfs(x-1, y)
		}
		if x+1 < m && board[x+1][y] == 'O' {
			dfs(x+1, y)
		}
		if y > 0 && board[x][y-1] == 'O' {
			dfs(x, y-1)
		}
		if y+1 < n && board[x][y+1] == 'O' {
			dfs(x, y+1)
		}
	}
	for _, p := range pos {
		dfs(p[0], p[1])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
