package is_graph_bipartite

import "testing"

func TestIGB(t *testing.T) {
	g := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}
	t.Log(isBipartite(g))
}
