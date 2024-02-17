package magical_string

func magicalString(n int) int {
	if n <= 1 {
		return n
	}

	var s []byte
	s = append(s, 1)
	var r int
	r = 1
	var j, k int
	for j < n-1 {
		if s[k] == 1 {
			if s[j] == 2 {
				r++
			}
			s = append(s, getRevert(s[j]))
			j++
			k++
		} else {
			r++
			s = append(s, s[j])
			s = append(s, getRevert(s[j]))
			j += 2
			k++
		}
	}
	if j == n && s[j] == 1 {
		r--
	}

	return r
}

func getRevert(x byte) byte {
	if x == 1 {
		return 2
	}
	return 1
}
