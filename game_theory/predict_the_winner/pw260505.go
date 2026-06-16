package predict_the_winner

func predictTheWinner(nums []int) bool {
	n := len(nums)
	var sum int
	prefix := make([]int, n+1)
	for i, num := range nums {
		sum += num
		prefix[i+1] = prefix[i] + num
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			dp[i][j] = maxInt(
				prefix[j+1]-prefix[i]-dp[i+1][j],
				prefix[j+1]-prefix[i]-dp[i][j-1],
			)
		}
	}

	if dp[0][n-1] >= (sum+1)/2 {
		return true
	}

	return false
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
