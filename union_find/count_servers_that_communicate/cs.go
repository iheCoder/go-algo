package count_servers_that_communicate

type unionFinder struct {
	parents []int
	rank    []int
}

func newUnionFinder(n int) *unionFinder {
	uf := &unionFinder{
		parents: make([]int, n),
		rank:    make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}

	return uf
}

func (uf *unionFinder) find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.find(uf.parents[x])
	}

	return uf.parents[x]
}

func (uf *unionFinder) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX == rootY {
		return
	}

	switch {
	case uf.rank[rootX] > uf.rank[rootY]:
		uf.parents[rootY] = rootX
	case uf.rank[rootY] > uf.rank[rootX]:
		uf.parents[rootX] = rootY
	default:
		uf.parents[rootY] = rootX
		uf.rank[rootX]++
	}
}

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := newUnionFinder(m * n)

	for i := 0; i < m; i++ {
		j := 0
		for j < n && grid[i][j] == 0 {
			j++
		}
		vn := i*n + j

		for ; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			uf.union(vn, i*n+j)
		}
	}

	for j := 0; j < n; j++ {
		i := 0
		for i < m && grid[i][j] == 0 {
			i++
		}
		vn := i*n + j

		for ; i < m; i++ {
			if grid[i][j] == 0 {
				continue
			}

			uf.union(vn, i*n+j)
		}
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && uf.rank[uf.find(i*n+j)] > 0 {
				ans++
			}
		}
	}

	return ans
}
