package check_if_there_is_a_valid_path_in_a_grid

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return true
	}

	const (
		U = 1 << 0
		D = 1 << 1
		L = 1 << 2
		R = 1 << 3
	)
	// 类型对应的方向
	masks := map[int]int{
		1: L | R,
		2: U | D,
		3: L | D,
		4: R | D,
		5: L | U,
		6: R | U,
	}

	dirs := []struct {
		dx, dy, bit, back int
	}{
		{-1, 0, U, D},
		{1, 0, D, U},
		{0, -1, L, R},
		{0, 1, R, L},
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	q := [][2]int{
		{0, 0},
	}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		s := masks[grid[x][y]]

		for _, dir := range dirs {
			if s&dir.bit == 0 {
				continue
			}

			nx, ny := x+dir.dx, y+dir.dy
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			if visited[[2]int{nx, ny}] {
				continue
			}

			ns := masks[grid[nx][ny]]
			if ns&dir.back == 0 {
				continue
			}

			visited[[2]int{nx, ny}] = true
			if nx == m-1 && ny == n-1 {
				return true
			}
			q = append(q, [2]int{nx, ny})
		}
	}

	return false
}
