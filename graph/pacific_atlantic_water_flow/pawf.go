package pacific_atlantic_water_flow

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	ps, as := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		ps[i] = make([]bool, n)
		as[i] = make([]bool, n)
	}

	var dfs func(i, j int, ocean [][]bool)
	dfs = func(i, j int, ocean [][]bool) {
		if ocean[i][j] {
			return
		}

		ocean[i][j] = true
		var ni, nj int
		for _, dir := range dirs {
			ni, nj = i+dir[0], j+dir[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && heights[ni][nj] >= heights[i][j] {
				dfs(ni, nj, ocean)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, ps)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, ps)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, as)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, as)
	}

	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ps[i][j] && as[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// 一个比较优雅的解决思路，可惜就是正在遍历的也会影响递归遍历的
// 	// 0代表都不通-1不知道1p通10a通11都通
//	marks[0][n-1], marks[m-1][0] = 11, 11
// func pacificAtlantic(heights [][]int) [][]int {
//	m, n := len(heights), len(heights[0])
//	marks := make([][]int, m)
//	for i := 0; i < m; i++ {
//		marks[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			marks[i][j] = -1
//		}
//	}
//	// 0代表都不通-1不知道1p通10a通11都通
//	marks[0][n-1], marks[m-1][0] = 11, 11
//
//	var dfs func(i, j int) int
//	dfs = func(i, j int) int {
//		if marks[i][j] >= 0 {
//			return marks[i][j]
//		}
//
//		marks[i][j] = 0
//		h := heights[i][j]
//		l, r := i == 0 || j == 0, i == m-1 || j == n-1
//		var v1, v2, v3, v4 int
//		if i > 0 && h >= heights[i-1][j] {
//			v1 = dfs(i-1, j)
//		}
//		if i < m-1 && h >= heights[i+1][j] {
//			v2 = dfs(i+1, j)
//		}
//		if j > 0 && h >= heights[i][j-1] {
//			v3 = dfs(i, j-1)
//		}
//		if j < n-1 && h >= heights[i][j+1] {
//			v4 = dfs(i, j+1)
//		}
//
//		sum := v1 + v2 + v3 + v4
//		if !l && sum%10 > 0 {
//			l = true
//		}
//		if !r && sum >= 10 {
//			r = true
//		}
//
//		switch {
//		case l && r:
//			marks[i][j] = 11
//		case l && !r:
//			marks[i][j] = 1
//		case !l && r:
//			marks[i][j] = 10
//		default:
//			marks[i][j] = 0
//		}
//
//		return marks[i][j]
//	}
//
//	var ans [][]int
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if dfs(i, j) > 10 {
//				ans = append(ans, []int{i, j})
//			}
//		}
//	}
//	return ans
//}
