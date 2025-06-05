package queens_that_can_attack_the_king

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	m := make(map[[2]int]int)
	for _, queen := range queens {
		m[[2]int{queen[0], queen[1]}] = -1
	}

	for _, dir := range dirs {
		for step := 1; step < 8; step++ {
			dx, dy := dir[0]*step, dir[1]*step
			i, j := king[0]+dx, king[1]+dy
			if i < 0 || i >= 8 || j < 0 || j >= 8 {
				break
			}

			k := [2]int{i, j}
			if m[k] < 0 {
				m[k] = 1
				break
			}
		}
	}

	ans := make([][]int, 0)
	for k, a := range m {
		if a > 0 {
			ans = append(ans, []int{k[0], k[1]})
		}
	}

	return ans
}
