package check_if_word_is_valid_after_substitutions

func isValid(s string) bool {
	stk := make([]byte, 0, len(s))
	checkIfEliminate := func() bool {
		n := len(stk)
		if n < 2 {
			return false
		}
		return stk[n-1] == 'b' && stk[n-2] == 'a'
	}
	for i := 0; i < len(s); i++ {
		if s[i] != 'c' {
			stk = append(stk, s[i])
			continue
		}

		if !checkIfEliminate() {
			return false
		}

		stk = stk[:len(stk)-2]
	}

	return len(stk) == 0
}

// 		for checkIfEliminate() {
//			stk = stk[:len(stk)-2]
//			if len(stk) < 3 || stk[len(stk)-1] != 'c' {
//				break
//			}
//		}
