package reverse_words_in_a_string

import (
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	sa := make([]string, 1)
	var m int
	var isSpace bool
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if !isSpace && i != len(s) {
				sa = append(sa, "")
				m++
			}
			isSpace = true
		} else {
			sa[m] += string(s[i])
			isSpace = false
		}
	}

	l, r := 0, len(sa)-1
	for l < r {
		sa[l], sa[r] = sa[r], sa[l]
		l++
		r--
	}

	return strings.Join(sa, " ")
}

// 	bs := make([]byte, 0)
//	var isSpace bool
//	for i := 0; i <= len(s); i++ {
//		if i == len(s) || s[i] == ' ' {
//			if !isSpace && i != len(s) {
//				bs = append(bs, ' ')
//			}
//			isSpace = true
//		} else {
//			isSpace = false
//			bs = append(bs, s[i])
//		}
//	}
//
//	j, k := 0, len(bs)-1
//	for j < k {
//		bs[j], bs[k] = bs[k], bs[j]
//		j++
//		k--
//	}
