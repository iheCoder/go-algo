package binary_gap

func binaryGap(n int) int {
	var ans, i int
	last := -1
	for n > 0 {
		if n&1 > 0 {
			if last != -1 {
				ans = maxInt(ans, i-last)
			}
			last = i
		}
		n >>= 1
		i++
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
