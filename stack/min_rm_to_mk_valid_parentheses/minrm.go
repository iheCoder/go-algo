package min_rm_to_mk_valid_parentheses

func minRemoveToMakeValid(s string) string {
	type pair struct {
		isLeft bool
		index  int
	}

	stk := make([]pair, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stk = append(stk, pair{
				isLeft: true,
				index:  i,
			})
		case ')':
			if len(stk) == 0 || !stk[len(stk)-1].isLeft {
				stk = append(stk, pair{
					isLeft: false,
					index:  i,
				})
				continue
			}

			stk = stk[:len(stk)-1]
		}
	}

	if len(stk) == 0 {
		return s
	}

	sb := make([]byte, 0, len(s))
	first := stk[0]
	for i := 0; i < len(s); i++ {
		if first.index == i {
			stk = stk[1:]
			if len(stk) > 0 {
				first = stk[0]
			}
		} else {
			sb = append(sb, s[i])
		}
	}

	return string(sb)
}
