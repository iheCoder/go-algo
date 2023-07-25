package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	rcm := make([]map[byte]bool, 18)
	for i := 0; i < len(rcm); i++ {
		rcm[i] = make(map[byte]bool)
	}
	gm := make([]map[byte]bool, 9)
	for i := 0; i < len(gm); i++ {
		gm[i] = make(map[byte]bool)
	}

	var item byte
	var gi int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			item = board[i][j]
			if item == '.' {
				continue
			}
			if rcm[i+9][item] {
				return false
			}
			if rcm[j][item] {
				return false
			}
			gi = (i/3)*3 + (j / 3)
			if gm[gi][item] {
				return false
			}
			rcm[i+9][item] = true
			rcm[j][item] = true
			gm[gi][item] = true
		}
	}
	return true
}
