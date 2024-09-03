package coloring_a_border

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	q := [][]int{
		{row, col},
	}
	oc := grid[row][col]
	borders := make([][]int, 0)

	visited[row][col] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y := p[0], p[1]
		isBorder := false
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != oc {
				isBorder = true
			} else if !visited[nx][ny] {
				visited[nx][ny] = true
				q = append(q, []int{nx, ny})
			}
		}

		if isBorder {
			borders = append(borders, p)
		}
	}

	for _, border := range borders {
		grid[border[0]][border[1]] = color
	}

	return grid
}
