package letter_case_permutation

import "unicode"

func letterCasePermutation(s string) []string {
	if len(s) == 0 {
		return nil
	}

	var idx []int
	var bs []rune
	for i, b := range s {
		if unicode.IsLetter(b) {
			bs = append(bs, unicode.ToLower(b))
			idx = append(idx, i)
		}
	}

	n := len(bs)
	var r [][]byte
	for mask := 0; mask < 1<<n; mask++ {
		ri := make([]byte, 0, n)
		for i := 0; i < n; i++ {
			if mask>>i&1 > 0 {
				ri = append(ri, byte(unicode.ToUpper(bs[i])))
			} else {
				ri = append(ri, byte(bs[i]))
			}
		}
		r = append(r, ri)
	}

	var ss []string
	sbs := []byte(s)
	for _, ri := range r {
		si := make([]byte, len(s))
		copy(si, sbs)
		for i, j := range idx {
			si[j] = ri[i]
		}
		ss = append(ss, string(si))
	}
	return ss
}
