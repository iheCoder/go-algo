package num_of_provinces

type UnionFind struct {
	// 储存节点父节点
	parent []int
	// 储存节点所在树深度
	rank []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

// Find 找到元素x所属根节点
func (uf *UnionFind) Find(x int) int {
	// 路径压缩。若当前节点父节点不是自己，也即不是根节点
	// 那么递归父节点直到找到根节点，并且将该路径上所有节点设置为根节点
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	// 返回找到的父节点
	return uf.parent[x]
}

// Union 合并x、y所属树
func (uf *UnionFind) Union(x, y int) {
	// 找到x、y所属根节点
	rootX, rootY := uf.Find(x), uf.Find(y)

	// 若x、y根节点一致，说明已经属于同树，返回
	if rootX == rootY {
		return
	}

	// 将深度更低的连接到深度更高的树，防止发生退化
	// 若深度一致，则任选一树，并将该树深度加一
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isConnected[i][j] > 0 {
				uf.Union(i, j)
			}
		}
	}

	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[uf.Find(i)] = true
	}
	return len(m)
}
