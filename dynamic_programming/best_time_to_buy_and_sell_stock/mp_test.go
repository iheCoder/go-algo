package best_time_to_buy_and_sell_stock

import "testing"

func TestMP(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit(prices))
}
