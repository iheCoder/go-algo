package ambiguous_coordinates

func ambiguousCoordinates(s string) []string {
	var ans []string
	s = s[1 : len(s)-1]
	for l := 1; l < len(s); l++ {
		x := getPart(s[:l])
		if len(x) == 0 {
			continue
		}
		y := getPart(s[l:])
		if len(y) == 0 {
			continue
		}
		for _, xi := range x {
			for _, yi := range y {
				ans = append(ans, "("+xi+", "+yi+")")
			}
		}
	}

	return ans
}

func getPart(s string) []string {
	var ans []string
	if s == "0" || s[0] != '0' {
		ans = append(ans, s)
	}
	if s[len(s)-1] == '0' {
		return ans
	}

	for i := 1; i < len(s); i++ {
		if i != 1 && s[0] == '0' {
			continue
		}
		ans = append(ans, s[:i]+"."+s[i:])
	}

	return ans
}

// 枚举还真比回溯法更好
// 	var f func(i, commaPos, pointPos int)
//	f = func(i, commaPos, pointPos int) {
//		if i == n {
//			x := make([]byte, 0, n+3)
//			for j := 0; j < n; j++ {
//				x = append(x, s[j])
//				if j == commaPos {
//					x = append(x, ',')
//				} else if j == pointPos {
//					x = append(x, '.')
//				}
//			}
//			ans = append(ans, string(x))
//			return
//		}
//
//		if commaPos < 0 {
//			f(i+1, i, pointPos)
//		}
//		if pointPos < 0 {
//			f(i+1, commaPos, i)
//		} else if commaPos > pointPos {
//			x := make([]byte, 0, n+3)
//			for j := 0; j <= i; j++ {
//				x = append(x, s[j])
//				if j == commaPos {
//					x = append(x, ',')
//				} else if j == pointPos {
//					x = append(x, '.')
//				}
//			}
//			x = append(x, '.')
//			for j := i + 1; j <= n; j++ {
//				x = append(x, s[j])
//			}
//			ans = append(ans, string(x))
//		}
//		f(i+1, commaPos, pointPos)
//	}
//
//	f(0, -1, -1)
