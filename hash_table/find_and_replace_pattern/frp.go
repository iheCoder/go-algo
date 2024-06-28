package find_and_replace_pattern

func findAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	n := len(pattern)
	pm := make(map[byte]bool)
	r := make([]bool, n)
	for i := 0; i < n; i++ {
		if pm[pattern[i]] {
			r[i] = true
			continue
		}
		pm[pattern[i]] = true
	}

	for _, word := range words {
		m := make(map[byte]byte)
		if len(word) != len(pattern) {
			continue
		}
		match := true
		for i := 0; i < n; i++ {
			v, ok := m[word[i]]
			if !ok {
				if r[i] {
					match = false
					break
				}
				m[word[i]] = pattern[i]
				continue
			}
			if v != pattern[i] {
				match = false
				break
			}
		}

		if match {
			ans = append(ans, word)
		}
	}

	return ans
}
