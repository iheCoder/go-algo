package can_mk_palindrome_from_substr

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	count := make([]int, n+1)
	// count[i]意味着以第i元素结尾的所有字母数量的奇偶
	for i := 1; i <= n; i++ {
		count[i] = count[i-1] ^ (1 << (s[i-1] - 'a'))
	}

	ans := make([]bool, 0, len(queries))
	for _, query := range queries {
		l, r, k := query[0], query[1], query[2]
		x := count[r+1] ^ count[l]

		var bits int
		for x > 0 {
			x &= x - 1
			bits++
		}

		ans = append(ans, bits <= 2*k+1)
	}

	return ans
}
