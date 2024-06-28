package rabbits_in_forest

import "sort"

func numRabbits(answers []int) int {
	sort.Ints(answers)
	n := len(answers)
	var ans int
	for i := n - 1; i >= 0; i-- {
		cnt := answers[i]
		ans += cnt + 1
		for cnt > 0 && i-1 >= 0 && answers[i] == answers[i-1] {
			cnt--
			i--
		}
	}

	return ans
}
