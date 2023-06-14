package best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	var ans int
	l := -1
	var diff int
	for i := 1; i < len(prices); i++ {
		diff = prices[i] - prices[i-1]
		if l == -1 && diff > 0 {
			l = i - 1
		} else if l != -1 && diff < 0 {
			ans += prices[i-1] - prices[l]
			l = -1
		}
	}
	if l != -1 && diff >= 0 {
		ans += prices[len(prices)-1] - prices[l]
	}
	return ans
}
