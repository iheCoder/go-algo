package best_sighseeing_pair

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]
	for i := 1; i < len(values); i++ {
		ans = maxInt(ans, mx+values[i]-i)
		mx = maxInt(mx, values[i]+i)
	}
	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
