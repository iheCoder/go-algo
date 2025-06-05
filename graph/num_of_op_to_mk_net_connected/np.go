package num_of_op_to_mk_net_connected

type unionFind struct {
	parent, size []int
	setCount     int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = i
	}

	return &unionFind{
		parent:   parent,
		size:     size,
		setCount: n,
	}
}

func (u *unionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}

	return u.parent[x]
}

func (u *unionFind) union(x, y int) {
	px, py := u.find(x), u.find(y)
	if px == py {
		return
	}

	if u.size[px] < u.size[py] {
		px, py = py, px
	}

	u.parent[py] = px
	u.size[px] += u.size[py]
	u.setCount--
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	uf := newUnionFind(n)
	for _, connection := range connections {
		uf.union(connection[0], connection[1])
	}

	return uf.setCount - 1
}
