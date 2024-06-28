package reorganize_string

func reorganizeString(s string) string {
	type pair struct {
		count int
		a     byte
	}

	m := make([]*pair, 26)
	for i := 0; i < 26; i++ {
		m[i] = &pair{
			a: byte('a' + i),
		}
	}

	n := len(s)
	var mx int
	for i := 0; i < n; i++ {
		j := s[i] - 'a'
		m[j].count++
		if m[j].count > mx {
			mx = m[j].count
		}
	}

	if mx > (n+1)/2 {
		return ""
	}

	ei, oi := 0, 1
	half := n / 2
	ans := make([]byte, n)
	for _, p := range m {
		for p.count > 0 && p.count <= half && oi < n {
			ans[oi] = p.a
			oi += 2
			p.count--
		}
		for p.count > 0 {
			ans[ei] = p.a
			p.count--
			ei += 2
		}
	}

	return string(ans)
}

// 显然这种双指针没啥用
// 	sort.Slice(m, func(i, j int) bool {
//		return m[i].count > m[j].count
//	})
//
//	i, j := 0, 1
//	ans := make([]byte, 0, n)
//	for {
//		if m[i].count > 0 {
//			ans = append(ans, m[i].a)
//			m[i].count--
//		}
//		for i < 26 && m[i].count <= 0 {
//			i++
//		}
//		if i >= 26 {
//			return string(ans)
//		}
//
//		if i >= j {
//			j = i + 1
//		}
//		if j < 26 && m[j].count > 0 {
//			ans = append(ans, m[j].a)
//			m[j].count--
//		}
//		for j < 26 && m[j].count <= 0 {
//			j++
//		}
//	}
