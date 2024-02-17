package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	type pair struct {
		temp  int
		index int
	}
	stk := make([]pair, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1].temp <= temperatures[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			answer[i] = stk[len(stk)-1].index - i
		}
		stk = append(stk, pair{temp: temperatures[i], index: i})
	}
	return answer
}
