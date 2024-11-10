package lexicographically_smallest_equivalent

import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := NewUnionFind(26)
	n := len(s1)
	for i := 0; i < n; i++ {
		uf.union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	var sb strings.Builder
	bl := len(baseStr)
	for i := 0; i < bl; i++ {
		root := uf.find(int(baseStr[i] - 'a'))
		sb.WriteByte(byte(root + 'a'))
	}

	return sb.String()
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent}
}

func (uf *UnionFind) union(x, y int) {
	rootx := uf.find(x)
	rooty := uf.find(y)
	if rootx != rooty {
		if rootx < rooty {
			uf.parent[rooty] = rootx
		} else {
			uf.parent[rootx] = rooty
		}
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
