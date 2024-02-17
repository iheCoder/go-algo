package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := len(s)
	hs := make(map[byte]int)
	for i := 0; i < n; i++ {
		hs[s[i]]++
	}

	for i := 0; i < n; i++ {
		if c, ok := hs[t[i]]; !ok || c <= 0 {
			return false
		}
		hs[t[i]]--
	}
	return true
}
