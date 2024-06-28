package knight_probability_in_chessboard

var dirs = [][]int{
	{1, 2},
	{2, 1},
	{2, -1},
	{1, -2},
	{-1, -2},
	{-2, -1},
	{-2, 1},
	{-1, 2},
}

func knightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1
	}
	if n <= 1 {
		return 0
	}

	m := make(map[int]float64)
	getKey := func(rem, i, j int) int {
		return rem*n*n + i*n + j
	}
	m[getKey(0, row, column)] = 1

	stk := make([][]int, 1)
	stk[0] = []int{row, column}
	var newStk [][]int
	for step := 1; step <= k; step++ {
		for len(stk) > 0 {
			x := stk[0]
			i, j := x[0], x[1]
			key := getKey(step-1, i, j)
			p := m[key]
			for _, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni < 0 || ni >= n || nj < 0 || nj >= n {
					continue
				}

				nkey := getKey(step, ni, nj)
				if _, ok := m[nkey]; !ok {
					newStk = append(newStk, []int{ni, nj})
				}
				m[nkey] += 0.125 * p
			}
			stk = stk[1:]
			delete(m, key)
		}
		stk = newStk
		newStk = [][]int{}
	}

	var ans float64
	for _, item := range stk {
		key := getKey(k, item[0], item[1])
		if v, ok := m[key]; ok {
			ans += v
			delete(m, key)
		}
	}
	return ans
}
