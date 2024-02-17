package compress_string_lcci

import "strconv"

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}

	a := S[0]
	c := 1
	sb := make([]byte, 0)
	for i := 1; i < len(S); i++ {
		if S[i] == a {
			c++
		} else {
			sb = append(sb, a)
			sb = append(sb, []byte(strconv.Itoa(c))...)
			a = S[i]
			c = 1
		}
	}
	sb = append(sb, a)
	sb = append(sb, []byte(strconv.Itoa(c))...)

	r := string(sb)
	if len(r) < len(S) {
		return r
	}
	return S
}
