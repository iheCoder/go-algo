package int_break

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = maxInt(curMax, maxInt(j*(i-j), dp[i-j]*j))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 以sqrt作为乘数个数是不行的了，以最高bit位也是不行
// 	if n == 2 {
//		return 1
//	}
//	if n == 3 {
//		return 2
//	}
//	c := int(math.Sqrt(float64(n)))
//	x := n / c
//	rem := n % x
//	y := 1
//	if rem != 0 {
//		c--
//		y = x + rem
//	}
//	return int(math.Pow(float64(x), float64(c))) * y
