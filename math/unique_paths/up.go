package unique_paths

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m > n {
		m, n = n, m
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[1][i] = 1
	}
	for i := 0; i <= m; i++ {
		dp[i][1] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

// 居然内存溢出了，离谱
func uniquePathsRecur(m int, n int) int {
	stkX, stkY := []byte{0}, []byte{0}
	var x, y byte
	vc := 1
	r := 1
	for len(stkX) != 0 {
		x, y = stkX[0], stkY[0]
		stkX, stkY = stkX[1:], stkY[1:]
		vc--

		if x < byte(m)-1 && y < byte(n)-1 {
			stkX = append(stkX, x)
			stkX = append(stkX, x+1)
			stkY = append(stkY, y+1)
			stkY = append(stkY, y)
		} else if x < byte(m)-1 {
			stkX = append(stkX, x+1)
			stkY = append(stkY, y)
		} else if y < byte(n)-1 {
			stkX = append(stkX, x)
			stkY = append(stkY, y+1)
		}

		if vc == 0 {
			vc = len(stkX)
			r = maxInt(r, vc)
		}
	}
	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
