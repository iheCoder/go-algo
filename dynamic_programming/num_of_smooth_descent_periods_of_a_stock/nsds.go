package num_of_smooth_descent_periods_of_a_stock

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	if n <= 1 {
		return int64(n)
	}

	ans, descent := int64(0), int64(0)
	cal := func(i int) {
		if descent <= 0 {
			ans++
			descent = 1
			return
		}

		last := prices[i-1]
		if last-prices[i] != 1 {
			ans++
			descent = 1
			return
		}

		descent++
		ans += descent
	}

	for i := 0; i < n; i++ {
		cal(i)
	}

	return ans
}
