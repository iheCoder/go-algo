package reverse_substrings_between_each_pair_of_parentheses

func reverseParentheses(s string) string {
	n := len(s)
	stk := make([]byte, 0, n)

	pop := func() []byte {
		tmp := make([]byte, 0)
		for len(stk) > 0 {
			last := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if last == '(' {
				break
			}
			tmp = append(tmp, last)
		}

		return tmp
	}

	for i := 0; i < n; i++ {
		if s[i] == ')' {
			stk = append(stk, pop()...)
		} else {
			stk = append(stk, s[i])
		}
	}

	return string(stk)
}
