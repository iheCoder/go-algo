package ugly_number_iii

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(a, bc)
	count := func(m int) int {
		return m/a + m/b + m/c - m/ab - m/ac - m/bc + m/abc
	}

	left, right := 1, int(2e9)
	for left < right {
		mid := left + (right-left)/2
		if count(mid) < n {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
