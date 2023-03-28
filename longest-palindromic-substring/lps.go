package longest_palindromic_substring

func longestPalindrome(s string) string {
	var longestStr string
	longestStr = string(s[0])
	li := 0
	ri := len(s) - 1
	for li < len(s) {
		for li <= ri {
			isPali := true
			// 长度应该还要在差加1
			if ri-li+1 <= len(longestStr) {
				break
			}

			if s[li] == s[ri] {
				i, j := li+1, ri-1
				for i <= j {
					if s[i] != s[j] {
						isPali = false
						break
					}
					i++
					j--
				}
				if isPali {
					ls := s[li : ri+1]
					if len(ls) > len(longestStr) {
						longestStr = ls
					}
					break
				}
			}

			ri--
		}
		li++
		ri = len(s) - 1
	}

	return longestStr
}
