package partition_arr_for_max_sum

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var maxV int
		for j := 1; j <= k && i-j >= 0; j++ {
			maxV = maxInt(maxV, arr[i-j])
			dp[i] = maxInt(dp[i], dp[i-j]+maxV*j)
		}
	}

	return dp[n]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
