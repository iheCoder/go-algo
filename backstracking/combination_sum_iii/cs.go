package combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	cur := make([]int, 0)
	m := make([]int, 10)
	var f func(rem, rk, i int)
	f = func(rem, rk, i int) {
		if rem == 0 && rk == 0 {
			a := make([]int, len(cur))
			copy(a, cur)
			ans = append(ans, a)
			return
		}
		if rk <= 0 {
			return
		}

		for ; i < 10; i++ {
			if m[i] > 0 {
				continue
			}
			if rem-i < 0 {
				return
			}
			m[i] = 1
			cur = append(cur, i)
			f(rem-i, rk-1, i+1)
			cur = cur[:len(cur)-1]
			m[i] = 0
		}
	}
	f(n, k, 1)
	return ans
}
