package smallest_string_with_swaps

import "sort"

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &unionFind{
		parent: parent,
		rank:   rank,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return
	}

	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px

	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := newUnionFind(n)

	// union circles
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}

	// group circles
	groups := make(map[int][]byte)
	for i := 0; i < n; i++ {
		p := uf.find(i)
		groups[p] = append(groups[p], s[i])
	}

	// sort group items
	for _, g := range groups {
		sort.Slice(g, func(x, y int) bool {
			return g[x] < g[y]
		})
	}

	result := make([]byte, n)
	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		p := uf.find(i)
		result[i] = groups[p][counts[p]]
		counts[p]++
	}

	return string(result)
}
