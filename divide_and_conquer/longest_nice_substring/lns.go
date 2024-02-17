package longest_nice_substring

func longestNiceSubstring(s string) string {
	n := len(s)
	for l := n; l > 0; l-- {
		for i := 0; i <= n-l; i++ {
			var diff int
			m := make(map[byte]bool)
			var j int
			for j = i; j < i+l; j++ {
				if !m[s[j]] {
					if !m[(findPeerByte(s[j]))] {
						diff++
					} else {
						diff--
					}
					m[s[j]] = true
				}
				if diff > i+l-j {
					break
				}
			}
			if diff == 0 {
				return s[i:j]
			}
		}
	}
	return ""
}

func findPeerByte(b byte) byte {
	if b >= 97 {
		return b - 32
	} else {
		return b + 32
	}
}
