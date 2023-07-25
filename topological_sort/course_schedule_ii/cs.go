package course_schedule_ii

type graph struct {
	count    int
	adj      [][]int
	inDegree []int
}

func NewGraph(count int) *graph {
	adj := make([][]int, count)
	for i := 0; i < count; i++ {
		adj[i] = make([]int, 0)
	}
	g := &graph{
		count:    count,
		adj:      adj,
		inDegree: make([]int, count),
	}
	return g
}

func (g *graph) addEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.inDegree[w]++
}

func (g *graph) TopoSort() []int {
	q := make([]int, 0, g.count)
	for v, d := range g.inDegree {
		if d == 0 {
			q = append(q, v)
		}
	}

	var v int
	var vid []int
	ans := make([]int, 0, g.count)
	for len(q) != 0 {
		v = q[0]
		q = q[1:]
		ans = append(ans, v)

		vid = g.adj[v]
		for _, w := range vid {
			g.inDegree[w]--
			if g.inDegree[w] == 0 {
				q = append(q, w)
			}
		}
	}

	if len(ans) >= g.count {
		return ans
	}
	return nil
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := NewGraph(numCourses)
	for _, p := range prerequisites {
		g.addEdge(p[1], p[0])
	}

	return g.TopoSort()
}
