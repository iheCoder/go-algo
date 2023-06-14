package sqrtx

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	l, r := 1, x>>1
	var mid, y int
	for l < r {
		mid = l + (r-l)>>1
		y = mid * mid
		if y < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if r*r > x {
		r--
	}
	return r
}
