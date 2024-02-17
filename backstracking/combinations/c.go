package combinations

func combine(n int, k int) [][]int {
	var ans [][]int
	var f func(c int, cur []int)
	f = func(c int, cur []int) {
		if len(cur) == k {
			ans = append(ans, cur)
			return
		}
		if len(cur)+n-c+1 < k {
			return
		}

		f(c+1, cur)
		curr := make([]int, len(cur))
		copy(curr, cur)
		curr = append(curr, c)
		f(c+1, curr)
	}
	f(1, make([]int, 0))
	return ans
}
