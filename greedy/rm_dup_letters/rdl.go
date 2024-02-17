package rm_dup_letters

func removeDuplicateLetters(s string) string {
	if len(s) <= 1 {
		return s
	}

	bcs := make([]int, 26)
	for i := 0; i < len(s); i++ {
		bcs[s[i]-'a']++
	}

	stk := make([]byte, 0, len(s))
	m := make(map[byte]bool)
	push := func(i int) {
		if m[s[i]] {
			bcs[s[i]-'a']--
			return
		}
		for len(stk) > 0 && s[i] <= stk[len(stk)-1] && bcs[stk[len(stk)-1]-'a'] > 1 {
			bcs[stk[len(stk)-1]-'a']--
			m[stk[len(stk)-1]] = false
			stk = stk[:len(stk)-1]
		}
		m[s[i]] = true
		stk = append(stk, s[i])
	}

	for i := 0; i < len(s); i++ {
		push(i)
	}

	return string(stk)
}

// 并不行，无法删除abca这种
// 	var rs []byte
//	ors := []byte(s)
//	var si int
//	for rc > 0 {
//		for i := 0; i < len(ors); i++ {
//			// 若当前字符不存在，则直接保留
//			// 若当前大于等于后面的字符且存在重复，那么就删去该字符
//			// 若当前小于后面字符，则保留
//			si = int(s[i] - 'a')
//			if bcs[si] <= 1 {
//				rs = append(rs, ors[i])
//				continue
//			}
//			if ors[i] < ors[i+1] {
//				rs = append(rs, ors[i])
//			} else {
//				rc--
//			}
//			bcs[si]--
//		}
//		ors = rs
//		rs = []byte{}
//	}
//	return string(rs)

//if s[i] < stk[0] {
//	for bcs[stk[len(stk)-1]] > 1 && stk[len(stk)-1] > s[i] {
//		bcs[stk[len(stk)-1]]--
//		stk = stk[:len(stk)-1]
//	}
//}

//
//for i := len(stk) - 1; i >= 0; i-- {
//	if bcs[stk[i]-'a'] > 1 {
//		bcs[stk[i]-'a']--
//		stk = append(stk[:i], stk[i+1:]...)
//	}
//}
