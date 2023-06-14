package number_of_islands

type unionFinder struct {
	count  int
	parent []int
	rank   []int
}

func NewUF(grid [][]byte) *unionFinder {
	count := 0
	m := len(grid)
	n := len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			}
			rank[i*n+j] = 0
		}
	}
	return &unionFinder{
		count:  count,
		parent: parent,
		rank:   rank,
	}
}

func (u *unionFinder) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	// 合并。若集合x集合数量大于等于y，则x合并y；小于，则y合并x；
	if rootX != rootY {
		if u.rank[rootX] > u.rank[rootY] {
			u.parent[rootY] = rootX
		} else if u.rank[rootX] < u.rank[rootY] {
			u.parent[rootX] = rootY
		} else {
			u.parent[rootY] = rootX
			u.rank[rootX]++
		}
		u.count--
	}
}

func (u *unionFinder) find(i int) int {
	// 若当前元素父不是自己，那么就向上递归找到父节点
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *unionFinder) getCount() int {
	return u.count
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	uf := NewUF(grid)
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				cur := i*nc + j
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.union(cur, i*nc+j-1)
				}
				if j+1 < nc && grid[i][j+1] == '1' {
					uf.union(cur, i*nc+j+1)
				}
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.union(cur, (i-1)*nc+j)
				}
				if i+1 < nr && grid[i+1][j] == '1' {
					uf.union(cur, (i+1)*nc+j)
				}
			}
		}
	}
	return uf.getCount()
}
