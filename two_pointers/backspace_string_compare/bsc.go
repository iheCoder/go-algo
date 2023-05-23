package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	var bi, bj int

	for {
		for i >= 0 {
			if s[i] == '#' {
				bi++
				i--
			} else if bi > 0 {
				i--
				bi--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				bj++
				j--
			} else if bj > 0 {
				j--
				bj--
			} else {
				break
			}
		}

		if j < 0 || i < 0 {
			if j < 0 && i < 0 {
				return true
			}
			return false
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
}

// 		for i >= 0 && s[i] == '#' {
//			bi++
//			i--
//		}
//		for bi > 0 {
//			i--
//			bi--
//		}
