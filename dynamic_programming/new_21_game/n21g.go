package new_21_game

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}

	dp := make([]float64, k+maxPts)
	for i := k; i <= n && i < k+maxPts; i++ {
		dp[i] = 1.0
	}

	dp[k-1] = 1.0 * float64(minInt(n-k+1, maxPts)) / float64(maxPts)
	for i := k - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - ((dp[i+maxPts+1] - dp[i+1]) / float64(maxPts))
	}

	return dp[0]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// exceed time complexity
// 	if k > n {
//		return 0
//	}
//
//	m := make(map[int]float64)
//
//	var f func(cur int) float64
//	f = func(cur int) float64 {
//		if cur > n {
//			return 0
//		}
//		if cur >= k {
//			return 1
//		}
//		if p, ok := m[cur]; ok {
//			fmt.Println(p)
//			return p
//		}
//
//		var total float64
//		for i := 1; i <= maxPts; i++ {
//			total += f(cur + i)
//		}
//
//		p := total / float64(maxPts)
//		m[cur] = p
//		return p
//	}
//
//	return f(0)
