package battleships_in_a_board

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	var ans int

	var count int
	for i := 0; i < m; i++ {
		count = 0
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				count++
				if count == 2 {
					ans++
				}
			} else {
				count = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		count = 0
		for i := 0; i < m; i++ {
			if board[i][j] == 'X' {
				count++
				if count == 2 {
					ans++
				}
			} else {
				count = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if i < m-1 && board[i+1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				if j < n-1 && board[i][j+1] == 'X' {
					continue
				}
				ans++
			}
		}
	}

	return ans
}
