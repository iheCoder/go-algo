package loud_and_rich

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	poorer := make([][]int, n)
	for _, ri := range richer {
		poorer[ri[1]] = append(poorer[ri[1]], ri[0])
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if ans[i] != -1 {
			return ans[i]
		}

		mx := quiet[i]
		j := i
		for _, p := range poorer[i] {
			k := dfs(p)
			if quiet[k] < mx {
				mx = quiet[k]
				j = k
			}
		}

		ans[i] = j
		return j
	}

	for i := 0; i < n; i++ {
		if ans[i] != -1 {
			continue
		}
		ans[i] = dfs(i)
	}
	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
