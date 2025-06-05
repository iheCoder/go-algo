package num_of_closed_islands

type unionFinder struct {
	parents []int
	rank    []int
}

func newUnionFind(n int) *unionFinder {
	uf := &unionFinder{
		parents: make([]int, n),
		rank:    make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
	return uf
}

func (u *unionFinder) find(x int) int {
	if u.parents[x] != x {
		u.parents[x] = u.find(u.parents[x])
	}

	return u.parents[x]
}

func (u *unionFinder) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}

	switch {
	case u.rank[xp] > u.rank[yp]:
		u.parents[yp] = xp
	case u.rank[yp] > u.rank[xp]:
		u.parents[xp] = yp
	default:
		u.parents[yp] = xp
		u.rank[xp]++
	}
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := newUnionFind(m * n)

	gen := func(x, y int) int {
		return x*n + y
	}

	vn := 0
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			uf.union(i*n, vn)
		}
		if grid[i][n-1] == 0 {
			uf.union(i*n+n-1, vn)
		}
	}
	for j := 1; j < n-1; j++ {
		if grid[0][j] == 0 {
			uf.union(j, vn)
		}
		if grid[m-1][j] == 0 {
			uf.union((m-1)*n+j, vn)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			if i > 0 && grid[i-1][j] == 0 {
				uf.union(gen(i, j), gen(i-1, j))
			}
			if j > 0 && grid[i][j-1] == 0 {
				uf.union(gen(i, j), gen(i, j-1))
			}
		}
	}

	cnt := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			cnt[uf.find(gen(i, j))] = true
		}
	}

	if cnt[uf.find(vn)] {
		delete(cnt, uf.find(vn))
	}
	return len(cnt)
}
