package camelcase_match

func camelMatch(queries []string, pattern string) []bool {
	r := make([]bool, len(queries))
	upperCache := make([]bool, len(pattern))
	for i := 0; i < len(pattern); i++ {
		upperCache[i] = isUpperCase(pattern[i])
	}

	isSkip := false
	var k int
	for i := 0; i < len(queries); i++ {
		k = 0
		isSkip = false
		for j := 0; j < len(pattern); j++ {
			if upperCache[j] {
				for ; k < len(queries[i]); k++ {
					if isUpperCase(queries[i][k]) {
						if queries[i][k] != pattern[j] {
							isSkip = true
						}
						break
					}
				}
			} else {
				for ; k < len(queries[i]); k++ {
					if queries[i][k] == pattern[j] {
						break
					}
					if isUpperCase(queries[i][k]) {
						isSkip = true
						break
					}
				}
			}
			if isSkip || k >= len(queries[i]) {
				isSkip = true
				break
			}
			// k 用过了应该++
			k++
		}
		for k++; k < len(queries[i]); k++ {
			if isUpperCase(queries[i][k]) {
				isSkip = true
				break
			}
		}
		if isSkip {
			continue
		}
		r[i] = true
	}

	return r
}

func isUpperCase(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
