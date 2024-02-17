package unique_substr_in_wraparound_str

func findSubstringInWraproundString(s string) int {
	n := len(s)
	dp := make([]int, 26)
	var k int
	for i := 0; i < n; i++ {
		if i > 0 && checkContinuous(s[i-1], s[i]) {
			k++
		} else {
			k = 1
		}
		dp[s[i]-'a'] = maxInt(dp[s[i]-'a'], k)
	}

	var sum int
	for _, p := range dp {
		sum += p
	}
	return sum
}

func checkContinuous(x, y byte) bool {
	return y-x == 1 || x-y == 25
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 如果出现zab...zab这种情况就多了
// 	encounter := make([]int, 26)
//	encounter[s[0]-'a'] = 1
//	sum := 1
//	for i := 1; i < n; i++ {
//		if checkContinuous(s[i-1], s[i]) {
//			sum += i + 1
//		} else {
//			sum++
//		}
//		if encounter[s[i]-'a'] > 0 {
//			sum--
//			continue
//		}
//		encounter[s[i]-'a'] = 1
//	}
