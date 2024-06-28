package valid_parenthesis_string

func checkValidString(s string) bool {
	n := len(s)
	stk := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			if len(stk) > 0 {
				li := len(stk) - 1
				if stk[li] == '*' {
					for j := li - 1; j >= 0; j-- {
						if stk[j] == '(' {
							stk[j], stk[li] = stk[li], stk[j]
							break
						}
					}
				}
				stk = stk[:len(stk)-1]
			} else {
				return false
			}
		} else {
			stk = append(stk, s[i])
		}
	}

	if len(stk) == 0 {
		return true
	}

	// (右边的星一定要大于等(数量
	var sc int
	for j := len(stk) - 1; j >= 0; j-- {
		if stk[j] == '*' {
			sc++
		} else {
			if sc <= 0 {
				return false
			}
			sc--
		}
	}

	return true
}
