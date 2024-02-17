package permutations_ii

func permuteUnique(nums []int) [][]int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	n := len(nums)

	var ans [][]int
	var f func(cur []int)
	f = func(cur []int) {
		if len(cur) == n {
			ans = append(ans, cur)
			return
		}

		for k, v := range m {
			if v <= 0 {
				continue
			}

			m[k]--
			curr := make([]int, len(cur))
			copy(curr, cur)
			curr = append(curr, k)
			f(curr)
			m[k]++
		}
	}
	f(make([]int, 0))
	return ans
}
