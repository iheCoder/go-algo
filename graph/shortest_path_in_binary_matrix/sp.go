package shortest_path_in_binary_matrix

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	type trace struct {
		i, j  int
		steps int
	}
	q := []trace{
		{
			steps: 1,
		},
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.i == n-1 && p.j == n-1 {
			return p.steps
		}

		for _, dir := range dirs {
			ni, nj := p.i+dir[0], p.j+dir[1]
			if ni < 0 || nj < 0 || ni >= n || nj >= n || grid[ni][nj] == 1 {
				continue
			}

			if visited[[2]int{ni, nj}] {
				continue
			}

			q = append(q, trace{
				i:     ni,
				j:     nj,
				steps: p.steps + 1,
			})

			visited[[2]int{ni, nj}] = true
		}
	}

	return -1
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
