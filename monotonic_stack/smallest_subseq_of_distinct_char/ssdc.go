package smallest_subseq_of_distinct_char

func smallestSubsequence(s string) string {
	if len(s) <= 1 {
		return s
	}

	n := len(s)
	bcs := make([]int, 26)
	for i := 0; i < n; i++ {
		bcs[s[i]-'a']++
	}

	em := make(map[byte]bool)
	stk := make([]byte, 0, n)

	push := func(i int) {
		if em[s[i]] {
			bcs[s[i]-'a']--
			return
		}

		for len(stk) > 0 && s[i] < stk[len(stk)-1] && bcs[stk[len(stk)-1]-'a'] > 1 {
			bcs[stk[len(stk)-1]-'a']--
			em[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}

		em[s[i]] = true
		stk = append(stk, s[i])
	}

	for i := 0; i < n; i++ {
		push(i)
	}

	return string(stk)
}
