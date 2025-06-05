package replace_the_substring_for_balanced_string

func balancedString(s string) int {
	cnts := make(map[byte]int)
	n := len(s)
	part := n / 4
	for i := 0; i < n; i++ {
		cnts[s[i]]++
	}

	check := func() bool {
		return cnts['Q'] <= part && cnts['W'] <= part && cnts['E'] <= part && cnts['R'] <= part
	}
	if check() {
		return 0
	}

	ans := n
	var r int
	for l := 0; l < n; l++ {
		for r < n && !check() {
			cnts[s[r]]--
			r++
		}

		if !check() {
			return ans
		}

		ans = minInt(ans, r-l)
		cnts[s[l]]++
	}

	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
