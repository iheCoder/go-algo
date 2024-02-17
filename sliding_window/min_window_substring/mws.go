package min_window_substring

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	// 0. 初始化，获取t中所有字母出现次数
	m := make(map[byte]int)
	var diff int
	for i := 0; i < len(t); i++ {
		m[t[i]]++
		diff++
	}
	// 1. 遍历一直找到下一个蕴含的值，差别减少，并记录所在位置
	cnts := make(map[byte]int)
	indexes := make([]int, 0)
	var i, ix int
	var r string
	for j := 0; j < len(s); j++ {
		if k, ok := m[s[j]]; ok {
			cnts[s[j]]++
			indexes = append(indexes, j)
			// 若仅在==k，diff--，那么一个字符有多个的时候就不对了
			if cnts[s[j]] <= k {
				diff--
			}
		}
		if diff == 0 {
			lastByte := s[indexes[ix]]
			// 删除根本不匹配的那些字符
			i = indexes[ix]
			// 删除匹配的多余字符
			for cnts[lastByte] > m[lastByte] {
				cnts[lastByte]--
				ix++
				i = indexes[ix]
				lastByte = s[i]
			}
			r = minStr(r, s[i:j+1])
			if len(r) == len(t) {
				return r
			}
			cnts[lastByte]--
			ix++
			if ix < len(indexes) {
				i = indexes[ix]
			} else {
				i++
			}
			diff++
		}
	}
	// 2. 当差别为0时记录此时的长度，根据之前记录的所在位置，找到第二个匹配的值，查看此时差别是否为0
	// 随后接着遍历直到遍历完成
	return r
}

func minStr(s1, s2 string) string {
	if len(s1) == 0 {
		return s2
	}
	if len(s1) < len(s2) {
		return s1
	}
	return s2
}
