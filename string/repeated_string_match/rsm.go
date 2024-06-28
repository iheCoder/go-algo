package repeated_string_match

func repeatedStringMatch(a string, b string) int {
	na, nb := len(a), len(b)
	var i, j int
	for j = 0; j < nb; j++ {
		if b[j] == a[0] {
			break
		}
	}
	if j >= nb {
		return -1
	}

	var c int
	if j > 0 {
		c = 1
	}
	for ; j < nb; j++ {
		if b[j] != a[i] {
			return -1
		}
		i++
		if i >= na {
			i = 0
			c++
		}
	}
	if i > 0 {
		c++
	}
	return c
}
