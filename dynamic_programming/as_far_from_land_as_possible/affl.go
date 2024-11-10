package as_far_from_land_as_possible

type pair struct {
	x, y int
}

var dir = []pair{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	var queue []pair

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, pair{
					x: i,
					y: j,
				})
			}
		}
	}

	if len(queue) == n*n || len(queue) == 0 {
		return -1
	}

	ans := -1
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range dir {
			ni, nj := p.x+d.x, p.y+d.y
			if ni < 0 || nj < 0 || ni >= n || nj >= n || grid[ni][nj] != 0 {
				continue
			}

			grid[ni][nj] = grid[p.x][p.y] + 1
			ans = maxInt(ans, grid[ni][nj]-1)
			queue = append(queue, pair{
				x: ni,
				y: nj,
			})
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// 0 0 1 1 1
// 0 1 1 0 0
// 0 0 1 1 0
// 1 0 0 0 0

// 	var dfs func(x, y int)
//	dfs = func(x, y int) {
//		p := pair{
//			x: x,
//			y: y,
//		}
//		visited[p] = true
//
//		trace := math.MaxInt
//		for _, d := range dir {
//			ni, nj := x+d.x, y+d.y
//			if ni < 0 || nj < 0 || ni >= n || nj >= n {
//				continue
//			}
//			np := pair{
//				x: ni,
//				y: nj,
//			}
//			if visited[np] {
//				continue
//			}
//
//			if dp[ni][nj] == math.MaxInt {
//				dfs(ni, nj)
//			}
//			trace = minInt(trace, dp[ni][nj]+1)
//		}
//
//		dp[x][y] = trace
//		visited[p] = false
//	}

// 	visited := make(map[pair]bool)
//
//	bfs := func(x, y int) {
//		stk := []pair{
//			{x, y},
//		}
//
//		for len(stk) > 0 {
//			i, j := stk[0].x, stk[0].y
//			stk = stk[1:]
//
//			for _, d := range dir {
//				ni, nj := i+d.x, j+d.y
//				if ni < 0 || nj < 0 || ni >= n || nj >= n {
//					continue
//				}
//				np := pair{
//					x: ni,
//					y: nj,
//				}
//				if visited[np] {
//					continue
//				}
//				visited[np] = true
//
//				if grid[ni][nj] == 1 {
//					dp[i][j] = 1
//					continue
//				}
//
//				if dp[ni][nj] < math.MaxInt {
//					dp[i][j] = minInt(dp[i][j], dp[ni][nj]+1)
//					continue
//				}
//
//				stk = append(stk, np)
//			}
//		}
//
//		visited = make(map[pair]bool)
//	}
//
//	reGet := func(i, j int) {
//		for _, d := range dir {
//			ni, nj := i+d.x, j+d.y
//			if ni < 0 || nj < 0 || ni >= n || nj >= n {
//				continue
//			}
//
//			if dp[ni][nj] < math.MaxInt {
//				dp[i][j] = minInt(dp[i][j], dp[ni][nj]+1)
//			}
//		}
//	}
//
//	var ans int
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == 1 {
//				continue
//			}
//
//			bfs(i, j)
//			reGet(i, j)
//			ans = maxInt(ans, dp[i][j])
//		}
//	}
//
//	return ans
