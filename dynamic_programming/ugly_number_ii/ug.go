package ugly_number_ii

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var p2, p3, p5 int
	for i := 1; i < n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = minInt(x2, minInt(x3, x5))
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}

	return dp[n-1]
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
