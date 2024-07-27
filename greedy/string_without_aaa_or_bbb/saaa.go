package string_without_aaa_or_bbb

func strWithout3a3b(a int, b int) string {
	var ans string
	ac, bc := "a", "b"
	if a < b {
		a, b = b, a
		ac, bc = bc, ac
	}
	for a > 0 || b > 0 {
		if a > 0 {
			ans += ac
			a--
		}
		if a >= b*2 && a > 0 {
			ans += ac
			a--
		}

		if b > 0 {
			ans += bc
			b--
		}
	}

	return ans
}
