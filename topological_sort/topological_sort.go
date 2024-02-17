package topological_sort

type graph struct {
	count    int
	adj      [][]int
	inDegree []int
}

func NewGraph(v int) *graph {
	adj := make([][]int, v)
	for i := 0; i < v; i++ {
		adj[i] = make([]int, 0)
	}
	g := &graph{
		count: v,
		adj:   adj,
	}
	return g
}

func (g *graph) addEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.inDegree[w]++
}

func (g *graph) topologicalSort() bool {
	q := make([]int, 0)
	for i := 0; i < g.count; i++ {
		if g.inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var c int
	for len(q) != 0 {
		v := q[0]
		q = q[1:]

		c++

		beg := g.adj[v]
		for i := 0; i < len(beg); i++ {
			g.inDegree[beg[i]]--
			if g.inDegree[beg[i]] == 0 {
				q = append(q, beg[i])
			}
		}
	}

	return c >= g.count
}
