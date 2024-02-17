package minesweeper

const (
	unMine      = 'M'
	emptySquare = 'E'
	blank       = 'B'
	reMine      = 'X'
)

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == unMine {
		board[x][y] = reMine
		return board
	}

	m, n := len(board), len(board[0])
	var getAdjMineCount func(i, j int) int
	getAdjMineCount = func(x, y int) int {
		var mineCount int
		i0, j0 := x-1, y-1
		for i := i0; i < i0+3; i++ {
			if i < 0 || i >= m {
				continue
			}
			for j := j0; j < j0+3; j++ {
				if j < 0 || j >= n {
					continue
				}
				if board[i][j] == unMine {
					mineCount++
				}
			}
		}
		return mineCount
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		mineCount := getAdjMineCount(x, y)
		if mineCount > 0 {
			board[x][y] = byte('0' + mineCount)
			return
		}

		board[x][y] = blank

		i0, j0 := x-1, y-1
		for i := i0; i < i0+3; i++ {
			if i < 0 || i >= m {
				continue
			}
			for j := j0; j < j0+3; j++ {
				if j < 0 || j >= n {
					continue
				}
				if board[i][j] == emptySquare {
					dfs(i, j)
				}
			}
		}
	}
	dfs(x, y)
	return board
}

// 	if x > 0 {
//		adj = append(adj, []int{x - 1, y})
//		if board[x-1][y] == unMine {
//			mineCount++
//		}
//	}
//	if y > 0 {
//
//	}
//	if x < m-1 {
//
//	}
//	if y < n-1 {
//
//	}
