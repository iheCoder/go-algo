package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	m := prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = maxInt(dp[i-1], prices[i]-m)
		if prices[i] < m {
			m = prices[i]
		}
	}
	return dp[len(prices)-1]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
