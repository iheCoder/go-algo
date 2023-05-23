package longest_uncommon_subsequence_ii

import (
	"sort"
)

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j]) || (len(strs[i]) == len(strs[j]) && strs[i] < strs[j])
	})

	pass := true
	for i := len(strs) - 1; i >= 0; i-- {
		if i > 0 && strs[i] == strs[i-1] {
			continue
		}
		for j := i + 1; j < len(strs); j++ {
			if isMatch(strs[j], strs[i]) {
				pass = false
				break
			}
		}
		if pass {
			return len(strs[i])
		}
		pass = true
	}
	return -1
}

func isMatch(a, b string) bool {
	if len(a) == 0 {
		return false
	}

	var j int
	for i := 0; i < len(a) && j < len(b); i++ {
		if a[i] == b[j] {
			j++
		}
	}
	return j == len(b)
}

// 	ms := strs[len(strs)-1]
//	r := len(ms)
//	var i int
//	for i = len(strs) - 2; i >= 0 && len(strs[i]) == r; i-- {
//		if strs[i] != ms {
//			return r
//		}
//	}
//	if i < 0 {
//		return -1
//	}
//	if i == len(strs)-2 {
//		return r
//	}
//
//	var skipSameStr string
//	for ; i >= 0; i-- {
//		if strs[i] == skipSameStr {
//			continue
//		}
//		if isMatch(ms, strs[i]) {
//			continue
//		}
//		if i == 0 || (len(strs[i]) != len(strs[i-1]) || strs[i] != strs[i-1]) {
//			return len(strs[i])
//		}
//		skipSameStr = strs[i]
//	}
//	return -1
