package count_binary_substrings

func countBinarySubstrings(s string) int {
	if len(s) < 2 {
		return 0
	}

	b := s[0]
	var r int
	i := 1
	left, right := 1, 1
	for i < len(s) && s[i] == b {
		left++
		i++
	}
	if i >= len(s) {
		return 0
	}
	b = s[i]
	i++

	for {
		for i < len(s) && s[i] == b {
			right++
			i++
		}
		if i >= len(s) {
			r += minInt(left, right)
			return r
		}
		b = s[i]
		i++
		r += minInt(left, right)
		left = right
		right = 1
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
