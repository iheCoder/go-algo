package kth_symbol_in_grammar

func kthGrammar(n int, k int) int {
	for n >= 0 && 1<<n >= k {
		n--
	}

	var r int
	for k > 1 {
		k -= 1 << n
		for n >= 0 && 1<<n >= k {
			n--
		}
		r = 1 & (r ^ 1)
	}
	return r
}
