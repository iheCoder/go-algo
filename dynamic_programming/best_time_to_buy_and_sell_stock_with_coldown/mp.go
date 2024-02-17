package best_time_to_buy_and_sell_stock_with_coldown

// 感觉这应该是二维dp，可是二维状态该如何定义呢？
// 对于i来说有（1 i-1冻结 （2 i-1空闲 （3 i-1卖出
// 空闲可以分为（1 无profit （2 刚过冻结。无论是哪种情况都是尝试买入i-1，卖出i
// 对于i-1卖出或者冻结情况，方程很简单就是dp[i]=dp[i-1]，但若i-1卖出，i就处于冻结；若i-1冻结，则i处于空闲；对于i-1空闲情况则是nums[i]-nums[i-1]+dp[i-1]<free>，只要在i卖出成功，i处于冻结状态
// 设0为空闲、1为冻结、2为卖出
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = -prices[0]
	// 0 持有股票 1 不持有，且下一个处于冻结 2 不持有，且下一个不处于冻结
	for i := 1; i < n; i++ {
		// i持有股票等于i-1持有股票和i-1不持有股票，且不处于冻结，在i买入的最大值
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][2]-prices[i])
		// i不持有股票，且处于冻结等于i-1持有股票，并在i卖出
		dp[i][1] = dp[i-1][0] + prices[i]
		// i不持有且不处于冻结等于i-1持有股票，在i卖出和i-1不持有且不处于冻结
		dp[i][2] = maxInt(dp[i-1][1], dp[i-1][2])
	}

	return maxInt(dp[n-1][1], dp[n-1][2])
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
