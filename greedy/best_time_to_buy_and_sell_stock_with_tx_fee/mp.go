package best_time_to_buy_and_sell_stock_with_tx_fee

// 如果下一个利润都不大于fee，那么就应该直接跳过去
func maxProfit(prices []int, fee int) int {
	profit := 0
	buy := prices[0] + fee
	for i := 1; i < len(prices); i++ {
		if buy > prices[i]+fee {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
