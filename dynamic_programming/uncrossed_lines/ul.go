package uncrossed_lines

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i, v := range nums1 {
		for j, w := range nums2 {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxInt(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
