package shortest_bridge

var steps = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

// 破纪录失败四次了！！！
//func shortestBridge(grid [][]int) int {
//	n := len(grid)
//
//	var found bool
//	var markDfs func(i, j int)
//	islans := make([][]int, 0)
//	visited := make(map[int]bool)
//	markDfs = func(i, j int) {
//		k := i*n + j
//		if visited[k] {
//			return
//		}
//		visited[k] = true
//		if grid[i][j] == 1 {
//			grid[i][j] = 2
//			islans = append(islans, []int{i, j})
//			found = true
//		}
//
//		isZero := grid[i][j] == 0
//		for _, step := range steps {
//			x, y := i+step[0], j+step[1]
//			if x < 0 || y < 0 || x >= n || y >= n {
//				continue
//			}
//
//			if grid[x][y] == 1 && (!found || !isZero) {
//				markDfs(x, y)
//			} else if grid[x][y] == 0 && !found {
//				markDfs(x, y)
//			}
//		}
//	}
//	markDfs(0, 0)
//
//	ans := n * n
//	var dfs func(i, j, count int, visited map[int]bool)
//	dfs = func(i, j, count int, vs map[int]bool) {
//		k := i*n + j
//		if vs[k] {
//			return
//		}
//		vs[k] = true
//		defer func() {
//			vs[k] = false
//		}()
//
//		for _, step := range steps {
//			x, y := i+step[0], j+step[1]
//			if x < 0 || y < 0 || x >= n || y >= n {
//				continue
//			}
//
//			if grid[x][y] == 1 {
//				ans = minInt(ans, count+1)
//				return
//			} else if grid[x][y] == 0 {
//				dfs(x, y, count+1, vs)
//			}
//		}
//	}
//
//	for _, islan := range islans {
//		visited = make(map[int]bool)
//		dfs(islan[0], islan[1], 0, visited)
//	}
//
//	return ans - 1
//}

func shortestBridge(grid [][]int) (step int) {
	type pair struct {
		x, y int
	}
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)

	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}

			q := []pair{}
			var dfs func(i, j int)
			dfs = func(i, j int) {
				grid[i][j] = -1
				q = append(q, pair{i, j})

				for _, d := range dirs {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						dfs(x, y)
					}
				}
			}
			dfs(i, j)

			for {
				tmp := q
				q = nil

				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return
							}

							if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}

				step++
			}
		}
	}
	return
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
