package largest_plus_sign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = n
		}
	}

	getKey := func(i, j int) int {
		return i*n + j
	}

	banned := make(map[int]bool)
	for _, mine := range mines {
		banned[getKey(mine[0], mine[1])] = true
	}

	var count int
	upDp := func(i, j int) {
		if banned[getKey(i, j)] {
			count = 0
		} else {
			count++
		}
		dp[i][j] = minInt(dp[i][j], count)
	}

	var ans int
	for i := 0; i < n; i++ {
		count = 0
		// left -> right
		for j := 0; j < n; j++ {
			upDp(i, j)
		}

		// right -> left
		count = 0
		for j := n - 1; j >= 0; j-- {
			upDp(i, j)
		}
	}

	for j := 0; j < n; j++ {
		// bottom -> top
		count = 0
		for i := 0; i < n; i++ {
			upDp(i, j)
		}
		count = 0
		for i := n - 1; i >= 0; i-- {
			upDp(i, j)
			ans = maxInt(ans, dp[i][j])
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
