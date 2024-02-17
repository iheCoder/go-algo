package lexicographical_numbers

func lexicalOrder(n int) []int {
	var dfs func(x int)
	res := make([]int, 0)
	dfs = func(x int) {
		if x > n {
			return
		}

		for i := 0; i < 10; i++ {
			y := x + i
			if y > n {
				return
			} else if y == 0 {
				continue
			}
			res = append(res, y)
			dfs(y * 10)
		}
	}

	dfs(0)
	return res
}
