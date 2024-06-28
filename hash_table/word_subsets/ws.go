package word_subsets

func wordSubsets(words1 []string, words2 []string) []string {
	maxCounts := [26]int{}
	for _, s := range words2 {
		tmpCounts := [26]int{}
		for i := 0; i < len(s); i++ {
			tmpCounts[s[i]-'a']++
		}

		for i := 0; i < 26; i++ {
			maxCounts[i] = maxInt(maxCounts[i], tmpCounts[i])
		}
	}

	var ans []string
	for _, word := range words1 {
		tmpCounts := [26]int{}
		for i := 0; i < len(word); i++ {
			tmpCounts[word[i]-'a']++
		}

		if satisfyCharArray(tmpCounts, maxCounts) {
			ans = append(ans, word)
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func satisfyCharArray(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

// 显然即使改成map也不行
//func wordSubsets(words1 []string, words2 []string) []string {
//	n := len(words1)
//	m1 := make([]map[byte]int, n)
//	for i := 0; i < n; i++ {
//		m1[i] = make(map[byte]int)
//		for j := 0; j < len(words1[i]); j++ {
//			m1[i][words1[i][j]]++
//		}
//	}
//
//	m2 := make([]map[byte]int, len(words2))
//	for i, s := range words2 {
//		m2[i] = make(map[byte]int)
//		for j := 0; j < len(s); j++ {
//			m2[i][s[j]]++
//		}
//	}
//
//	var ans []string
//	for i := 0; i < n; i++ {
//		uni := true
//		for _, cs := range m2 {
//			if !satisfyCharArray(m1[i], cs) {
//				uni = false
//				break
//			}
//		}
//
//		if uni {
//			ans = append(ans, words1[i])
//		}
//	}
//
//	return ans
//}

//func satisfyCharArray(a, b map[byte]int) bool {
//	if len(a) < len(b) {
//		return false
//	}
//
//	for c, count := range b {
//		if a[c] < count {
//			return false
//		}
//	}
//
//	return true
//}
