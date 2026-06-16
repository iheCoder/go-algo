package best_time_to_buy_and_sell_stock_iv

import "math"

func maxProfit(k int, prices []int) int {
	buys, sells := make([]int, k+1), make([]int, k+1)
	for i := 1; i <= k; i++ {
		buys[i], sells[i] = math.MinInt, math.MinInt
	}

	for _, price := range prices {
		for j := 1; j <= k; j++ {
			buys[j] = maxInt(buys[j], sells[j-1]-price)
			sells[j] = maxInt(sells[j], buys[j]+price)
		}
	}

	return sells[k]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
