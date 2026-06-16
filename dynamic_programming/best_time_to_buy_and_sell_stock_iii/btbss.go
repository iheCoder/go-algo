package best_time_to_buy_and_sell_stock_iii

import "math"

func maxProfit(prices []int) int {
	buy1, sell1, buy2, sell2 := math.MinInt, math.MinInt, math.MinInt, math.MinInt

	for i := 0; i < len(prices); i++ {
		buy1 = maxInt(buy1, -prices[i])
		sell1 = maxInt(sell1, prices[i]+buy1)

		buy2 = maxInt(buy2, sell1-prices[i])
		sell2 = maxInt(sell2, buy2+prices[i])
	}

	return sell2
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
