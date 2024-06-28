package is_graph_bipartite

func isBipartite(graph [][]int) bool {
	n := len(graph)
	m := make([]int, n)
	var start int
	for ; start < len(graph); start++ {
		if len(graph[start]) > 0 {
			break
		}
	}
	if start >= len(graph) {
		return true
	}
	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, gi := range graph[i] {
			if m[gi] != 0 {
				if m[gi] == m[i] {
					return false
				}
				continue
			}

			m[gi] = -m[i]
			if !dfs(gi) {
				return false
			}
		}

		return true
	}

	for i := start; i < len(graph); i++ {
		if m[i] != 0 {
			continue
		}
		m[i] = 1
		if !dfs(i) {
			return false
		}
	}

	return true
}
