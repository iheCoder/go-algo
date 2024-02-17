package interleaving_string

// 我明白为什么要动态规划了，因为碰到相同的不知道怎么选择
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	n, m := len(s1), len(s2)
	dp := make([]bool, m+1)
	// 空字符必定匹配
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// i+j-1直接覆盖s3的所有范围
			k := i + j - 1
			// 若当前字符与s1字符匹配，且上轮匹配，那么本轮匹配
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[k]
			}
			// 若当前字符与s2字符匹配，且上轮匹配，或者是上步匹配，那么本轮匹配
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[k]
			}
		}
	}
	return dp[m]
}

// 	var matched bool
//	var i, j, k int
//	m, n := len(s1), len(s2)
//	for (i < m || j < n) && k < len(s3) {
//		matched = false
//		if i < m && s1[i] == s3[k] {
//			matched = true
//			i++
//		} else if j < n && s2[j] == s3[k] {
//			matched = true
//			j++
//		}
//		if !matched {
//			return false
//		}
//		k++
//	}
//	return k == len(s3)
