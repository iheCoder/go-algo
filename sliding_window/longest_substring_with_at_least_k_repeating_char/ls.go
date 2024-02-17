package longest_substring_with_at_least_k_repeating_char

func longestSubstring(s string, k int) int {
	if k > len(s) {
		return 0
	}

	var ans int
	for t := 1; t <= 26; t++ {
		var lessK, l, total int
		cnts := make([]int, 26)
		for r, ch := range s {
			if cnts[ch-'a'] == 0 {
				lessK++
				total++
			}
			cnts[ch-'a']++
			if cnts[ch-'a'] == k {
				lessK--
			}

			for total > t {
				if cnts[s[l]-'a'] == k {
					lessK++
				}
				cnts[s[l]-'a']--
				if cnts[s[l]-'a'] == 0 {
					lessK--
					total--
				}
				l++
			}
			if lessK == 0 {
				ans = maxInt(ans, r-l+1)
			}
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

//m := make(map[byte]int)
//var tempMax int
//for i := 0; i < len(s); i++ {
//	m[s[i]]++
//	tempMax = maxInt(tempMax, m[s[i]])
//}
//if tempMax < k {
//	return 0
//}
//
//var r int
//var i, j int
//tm := make(map[byte]int)
//for {
//	if j < len(s) && m[s[j]] >= k {
//		tm[s[j]]++
//		j++
//	} else {
//		l := j - i
//		// "bbaaacbd"这种就不行了
//		pass := true
//		for b, c := range tm {
//			if c < k {
//				pass = false
//			}
//			m[b] -= c
//			tm[b] = 0
//		}
//		if pass {
//			r = maxInt(r, l)
//		}
//		if j >= len(s) {
//			break
//		}
//		j++
//		i = j
//	}
//}
//return r
