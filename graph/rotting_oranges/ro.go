package rotting_oranges

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func orangesRotting(grid [][]int) int {
	ros := make([][]int, 0)
	var freshCount int
	for i, g := range grid {
		for j, x := range g {
			if x == 1 {
				freshCount++
			} else if x == 2 {
				ros = append(ros, []int{i, j})
			}
		}
	}

	if freshCount == 0 {
		return 0
	}
	if len(ros) == 0 {
		return -1
	}

	m, n := len(grid), len(grid[0])
	travel := func(i, j int) {
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < 0 || nj < 0 || ni >= m || nj >= n {
				continue
			}
			if grid[ni][nj] == 1 {
				grid[ni][nj] = 2
				freshCount--
				ros = append(ros, []int{ni, nj})
			}
		}
	}

	var ans int
	oldFreshCount := freshCount
	for {
		for _, ro := range ros {
			travel(ro[0], ro[1])
		}

		ans++
		if freshCount == 0 {
			return ans
		}
		if oldFreshCount == freshCount {
			return -1
		}
		oldFreshCount = freshCount
	}
}
