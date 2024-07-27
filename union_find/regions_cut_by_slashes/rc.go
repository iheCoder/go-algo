package regions_cut_by_slashes

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		count:  n,
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
}

func (u *UnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}

	return u.parent[x]
}

func (u *UnionFind) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX != rootY {
		u.parent[rootX] = rootY
		u.count--
	}
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := NewUnionFind(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := 4 * (i*n + j)
			c := grid[i][j]
			switch c {
			case '/':
				uf.union(k, k+3)
				uf.union(k+1, k+2)

			case '\\':
				uf.union(k, k+1)
				uf.union(k+2, k+3)

			default:
				uf.union(k, k+1)
				uf.union(k+2, k+1)
				uf.union(k+2, k+3)
			}

			// 第二个格和右边的第四个格合并
			if j < n-1 {
				uf.union(k+1, 4*(i*n+j+1)+3)
			}

			// 第三个格和下边的第一个格合并
			if i < n-1 {
				uf.union(k+2, 4*((i+1)*n+j))
			}

		}
	}

	return uf.count
}
