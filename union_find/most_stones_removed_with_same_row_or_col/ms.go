package most_stones_removed_with_same_row_or_col

type UnionFind struct {
	parent map[int]int
	count  int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		count:  0,
	}
}

func (u *UnionFind) getCount() int {
	return u.count
}

func (u *UnionFind) find(x int) int {
	if _, ok := u.parent[x]; !ok {
		u.parent[x] = x
		u.count++
	}

	if x != u.parent[x] {
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

func removeStones(stones [][]int) int {
	uf := NewUnionFind()

	for _, stone := range stones {
		uf.union(stone[0]+10001, stone[1])
	}

	return len(stones) - uf.count
}
