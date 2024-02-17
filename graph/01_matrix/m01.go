package _1_matrix

var dirs = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	getKey := func(i, j int) int {
		return i*n + j
	}
	visited := make(map[int]bool)

	stk := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				stk = append(stk, []int{i, j})
				visited[getKey(i, j)] = true
			}
		}
	}

	var i, j int
	var key int
	for len(stk) > 0 {
		i, j = stk[0][0], stk[0][1]
		stk = stk[1:]
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			key = getKey(x, y)
			if visited[key] {
				continue
			}
			ans[x][y] = ans[i][j] + 1
			visited[key] = true
			stk = append(stk, []int{x, y})
		}
	}

	return ans
}

// 这种寻找最近并不适合dfs
//var dfs func(i, j int) int
//dfs = func(i, j int) int {
//	if mat[i][j] == 0 {
//		return 0
//	}
//	if ans[i][j] != 0 {
//		return ans[i][j]
//	}
//
//	distance := math.MaxInt
//	key := getKey(i, j)
//	visited[key] = true
//
//	for _, dir := range dirs {
//		x, y := i+dir[0], j+dir[1]
//		if x < 0 || y < 0 || x >= m || y >= n {
//			continue
//		}
//		key = getKey(x, y)
//		if visited[key] {
//			continue
//		}
//		d := dfs(x, y)
//		if d < distance-1 {
//			distance = d + 1
//		}
//	}
//	return distance
//}
//for i := 0; i < m; i++ {
//	for j := 0; j < n; j++ {
//		if mat[i][j] == 0 {
//			continue
//		}
//		ans[i][j] = dfs(i, j)
//		visited = map[int]bool{}
//	}
//}
