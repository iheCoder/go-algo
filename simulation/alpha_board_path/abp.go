package alpha_board_path

var m = map[byte][2]int{
	'a': {0, 0},
	'b': {0, 1},
	'c': {0, 2},
	'd': {0, 3},
	'e': {0, 4},
	'f': {1, 0},
	'g': {1, 1},
	'h': {1, 2},
	'i': {1, 3},
	'j': {1, 4},
	'k': {2, 0},
	'l': {2, 1},
	'm': {2, 2},
	'n': {2, 3},
	'o': {2, 4},
	'p': {3, 0},
	'q': {3, 1},
	'r': {3, 2},
	's': {3, 3},
	't': {3, 4},
	'u': {4, 0},
	'v': {4, 1},
	'w': {4, 2},
	'x': {4, 3},
	'y': {4, 4},
	'z': {5, 0},
}

func alphabetBoardPath(target string) string {
	var path string
	at := [2]int{0, 0}
	for i := 0; i < len(target); i++ {
		to := m[target[i]]
		ym, xm := to[0]-at[0], to[1]-at[1]
		if at != [2]int{5, 0} {
			if xm > 0 {
				path += nChar(xm, 'R')
			} else if xm < 0 {
				path += nChar(-xm, 'L')
			}

			if ym > 0 {
				path += nChar(ym, 'D')
			} else if ym < 0 {
				path += nChar(-ym, 'U')
			}
		} else {
			if ym > 0 {
				path += nChar(ym, 'D')
			} else if ym < 0 {
				path += nChar(-ym, 'U')
			}

			if xm > 0 {
				path += nChar(xm, 'R')
			} else if xm < 0 {
				path += nChar(-xm, 'L')
			}
		}

		// how to handle z ?

		path += "!"
		at = to
	}

	return path
}

func nChar(n int, c byte) string {
	var ret []byte
	for i := 0; i < n; i++ {
		ret = append(ret, c)
	}

	return string(ret)
}
