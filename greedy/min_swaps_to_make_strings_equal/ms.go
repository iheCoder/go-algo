package min_swaps_to_make_strings_equal

func minimumSwap(s1 string, s2 string) int {
	xy, yx := 0, 0
	n := len(s1)
	for i := 0; i < n; i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			xy++
		}
		if s1[i] == 'y' && s2[i] == 'x' {
			yx++
		}
	}
	if (xy+yx)%2 == 1 {
		return -1
	}

	return xy/2 + yx/2 + (xy%2 + yx%2)
}

// 	var xc1, xc int
//	n := len(s1)
//	for i := 0; i < n; i++ {
//		if s1[i] == 'x' {
//			xc1++
//			xc++
//		}
//		if s2[i] == 'x' {
//			xc++
//		}
//	}
//	if xc%2 == 1 {
//		return -1
//	}
//
//	ans := n / 2
//	if xc1 > 0 && xc1 < n {
//		ans += n - xc1
//	}
//	return ans
