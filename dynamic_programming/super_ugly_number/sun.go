package super_ugly_number

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	m := len(primes)
	dp := make([]int, n+1)
	dp[1] = 1
	ps := make([]int, m)
	for i := 0; i < m; i++ {
		ps[i] = 1
	}

	for i := 2; i <= n; i++ {
		x := math.MaxInt
		for j, p := range ps {
			x = minInt(x, dp[p]*primes[j])
		}
		dp[i] = x
		for j, p := range ps {
			if x == dp[p]*primes[j] {
				ps[j]++
			}
		}
	}
	return dp[n]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
