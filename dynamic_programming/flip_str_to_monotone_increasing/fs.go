package flip_str_to_monotone_increasing

func minFlipsMonoIncr(s string) int {
	dp0, dp1 := 0, 0
	for _, c := range s {
		d0, d1 := dp0, minInt(dp0, dp1)
		if c == '1' {
			d0++
		} else {
			d1++
		}
		dp0, dp1 = d0, minInt(d0, d1)
	}

	return minInt(dp0, dp1)
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 错了一种情况，那就是可以将排前面的1变为0，反而会更少的flip次数
//func minFlipsMonoIncr(s string) int {
//	n := len(s)
//	ans := n
//
//	var encounter1 bool
//	var zero2one int
//	for i := 0; i < n; i++ {
//		if encounter1 && s[i] == '0' {
//			zero2one++
//			continue
//		}
//		if s[i] == '1' {
//			encounter1 = true
//		}
//	}
//	ans = minInt(ans, zero2one)
//
//	var encounter0 bool
//	var one2zero int
//	for i := n - 1; i >= 0; i-- {
//		if encounter0 && s[i] == '1' {
//			one2zero++
//			continue
//		}
//		if s[i] == '0' {
//			encounter0 = true
//		}
//	}
//
//	return minInt(ans, one2zero)
//}
