package decoded_string_at_index

func decodeAtIndex(s string, k int) string {
	n := len(s)
	var total int
	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			total *= int(s[i] - '0')
		} else {
			total++
		}
	}

	for i := n - 1; i >= 0; i-- {
		k %= total
		if s[i] >= '0' && s[i] <= '9' {
			total /= int(s[i] - '0')
			continue
		}
		if k == 0 {
			return string(s[i])
		}

		total--
	}

	return ""
}

// 这种暴力法显然是不行的，得另寻他法了！
func decodeAtIndexForce(s string, k int) string {
	stk := make([]byte, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		if len(stk) >= k {
			return string(stk[k-1])
		}

		if '0' <= s[i] && s[i] <= '9' {
			d := int(s[i] - '0')
			if len(stk)*d >= k {
				return string(stk[(k-1)%len(stk)])
			}

			tmp := make([]byte, len(stk))
			copy(tmp, stk)
			for j := 0; j < d-1; j++ {
				stk = append(stk, tmp...)
			}
			continue
		}

		stk = append(stk, s[i])
	}

	return string(stk[k-1])
}
