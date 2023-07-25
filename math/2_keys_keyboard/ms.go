package __keys_keyboard

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	maxSteps := 1000
	for i := 2; i < n+1; i++ {
		dp[i] = maxSteps
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= i>>1; j++ {
			if i%j != 0 {
				continue
			}
			dp[i] = minInt(dp[i], minInt(dp[j]+i/j, dp[i/j]+j))
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

// 复制不仅仅只能粘贴一次，比如9就可以3粘贴2次得到9
// 总之奇数直接返回奇数是不对的
//func minSteps(n int) int {
//	if n == 1 {
//		return 0
//	}
//	if n&1 > 0 {
//		return n
//	}
//
//	var x int
//	for n > 1 {
//		if n&1 > 0 {
//			return x<<1 + n
//		}
//		x++
//		n = n >> 1
//	}
//	return x << 1
//}
