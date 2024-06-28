package possible_bipartition

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n+1)
	for _, dislike := range dislikes {
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}

	parts := make([]int, n+1)

	var dfs func(i, part int) bool
	dfs = func(i, part int) bool {
		if parts[i] != 0 {
			if parts[i] != part {
				return false
			}
			return true
		}
		parts[i] = part

		for _, v := range graph[i] {
			if !dfs(v, -part) {
				return false
			}
		}

		return true
	}

	for i := 0; i <= n; i++ {
		if parts[i] == 0 {
			if !dfs(i, 1) {
				return false
			}
		}
	}

	return true
}
