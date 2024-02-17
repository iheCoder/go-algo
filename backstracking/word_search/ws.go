package word_search

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	if row*col < len(word) {
		return false
	}
	n := len(word)
	var f func(x, y, ti int) bool
	f = func(x, y, ti int) bool {
		if ti >= n {
			return true
		}

		if word[ti] != board[x][y] {
			return false
		}
		if ti+1 >= n {
			return true
		}

		board[x][y] = ' '
		defer func() {
			board[x][y] = word[ti]
		}()
		if x+1 < row && f(x+1, y, ti+1) {
			return true
		}
		if y+1 < col && f(x, y+1, ti+1) {
			return true
		}
		if x-1 >= 0 && f(x-1, y, ti+1) {
			return true
		}
		if y-1 >= 0 && f(x, y-1, ti+1) {
			return true
		}
		return false
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == ' ' {
				continue
			}

			if f(i, j, 0) {
				return true
			}
		}
	}
	return false
}
