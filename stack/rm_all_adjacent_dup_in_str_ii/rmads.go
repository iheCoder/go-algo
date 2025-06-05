package rm_all_adjacent_dup_in_str_ii

func removeDuplicates(s string, k int) string {
	type pair struct {
		x byte
		c int
	}

	n := len(s)
	stk := make([]pair, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && len(stk) > 0 && stk[len(stk)-1].x == s[i] {
			stk[len(stk)-1].c++
		} else {
			stk = append(stk, pair{
				x: s[i],
				c: 1,
			})
		}

		if stk[len(stk)-1].c == k {
			stk = stk[:len(stk)-1]
		}
	}

	r := make([]byte, 0, n)
	for _, p := range stk {
		for p.c > 0 {
			r = append(r, p.x)
			p.c--
		}
	}

	return string(r)
}
