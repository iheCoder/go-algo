package course_schedule

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

func (g *graph) hasTopoSort() bool {
	q := make([]int, 0, g.count)
	for v, d := range g.inDegree {
		if d == 0 {
			q = append(q, v)
		}
	}

	var c, v int
	var vid []int
	for len(q) != 0 {
		v = q[0]
		q = q[1:]
		c++

		vid = g.adj[v]
		for _, w := range vid {
			g.inDegree[w]--
			if g.inDegree[w] == 0 {
				q = append(q, w)
			}
		}
	}

	return c >= g.count
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := NewGraph(numCourses)
	for _, p := range prerequisites {
		g.addEdge(p[0], p[1])
	}

	return g.hasTopoSort()
}
