package is_subsequence

func isSubsequenceNormal(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	var count int
	var i, j int
	for i = 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			count++
			j++
		}
	}

	return count == len(s)
}

// 有每次访问加缓存的想法
// 或者将t拆开得到所有组合，不符合组合的返回
// 可是前两者在t比较大的情况下，内存消耗大
//
func isSubsequenceEnormous(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	if len(s) > len(t) {
		return false
	}
	if len(s) == len(t) {
		return s == t
	}

	var count int
	var i, j int
	for i = 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			count++
			j++
		}
	}

	return count == len(s)
}
