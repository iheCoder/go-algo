package min_add_to_make_parentheses_valid

func minAddToMakeValid(s string) int {
	n := len(s)
	stk := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stk = append(stk, '(')
		} else {
			if len(stk) > 0 && stk[len(stk)-1] == '(' {
				stk = stk[:len(stk)-1]
			} else {
				stk = append(stk, ')')
			}
		}
	}

	return len(stk)
}
