package gen_parentheses

func generateParenthesis(n int) []string {
	// 先生产n个左括号，然后第一个括号位置可以1个右括号，第二个位置可以两个右括号，以此类推
	// 若前面都没有括号，此时想要括号，则要全部补齐括号
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '('
	}
	return gps(n, 1, 1, string(bs))
}

func gps(n, index, ava int, s string) []string {
	if n <= 0 {
		return []string{s}
	}
	if ava > n {
		return nil
	}
	// 选择不填
	// 选择填
	var tmp []byte
	r := make([]string, 0, ava)
	for i := 0; i <= ava; i++ {
		ns := s
		if index >= len(ns) {
			ns += string(tmp)
		} else {
			ns = ns[:index] + string(tmp) + ns[index:]
		}
		r = append(r, gps(n-i, index+i+1, ava-i+1, ns)...)
		tmp = append(tmp, ')')
	}
	return r
}

// 	if n <= 0 {
//		return []string{s}
//	}
//	if ava > n {
//		return nil
//	}
//	// 选择不填
//	notFill := gps(n, index+1, ava+1, s)
//	// 选择填
//	var tmp []byte
//	for i := 0; i < ava; i++ {
//		tmp = append(tmp, ')')
//	}
//	if index >= len(s) {
//		s += string(tmp)
//	} else {
//		s = s[:index] + string(tmp) + s[index:]
//	}
//	fill := gps(n-ava, index+len(tmp)+1, 1, s)
//	return append(notFill, fill...)
