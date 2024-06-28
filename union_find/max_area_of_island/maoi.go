package max_area_of_island

type unionFind struct {
	parents []int
	ranks   []int
}

func newUnionFind(size int) *unionFind {
	parents := make([]int, size)
	ranks := make([]int, size)
	for i := 0; i < size; i++ {
		parents[i] = i
		ranks[i] = 0
	}

	return &unionFind{
		parents: parents,
		ranks:   ranks,
	}
}

func (u *unionFind) find(x int) int {
	if u.parents[x] != x {
		p := u.find(u.parents[x])
		u.parents[x] = p
	}

	return u.parents[x]
}

func (u *unionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}

	if u.ranks[xp] >= u.ranks[yp] {
		u.parents[yp] = xp
		u.ranks[xp] += u.ranks[yp]
	} else if u.ranks[xp] < u.ranks[yp] {
		u.parents[xp] = yp
		u.ranks[yp] += u.ranks[xp]
	}
}

func (u *unionFind) getMaxRank() int {
	var mr int
	for _, rank := range u.ranks {
		if mr < rank {
			mr = rank
		}
	}
	return mr
}

var dirs = [][]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	uf := newUnionFind(m * n)
	getKey := func(i, j int) int {
		return i*n + j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			key := getKey(i, j)
			if uf.ranks[key] == 0 {
				uf.ranks[key] = 1
			}
			for _, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni < 0 || nj < 0 || ni >= m || nj >= n || grid[ni][nj] == 0 {
					continue
				}

				nkey := getKey(ni, nj)
				if uf.ranks[nkey] == 0 {
					uf.ranks[nkey] = 1
				}
				uf.union(key, nkey)
			}
		}
	}

	return uf.getMaxRank()
}
