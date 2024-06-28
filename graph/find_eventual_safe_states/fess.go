package find_eventual_safe_states

func eventualSafeNodes(graph [][]int) []int {
	visited := make([]int, len(graph))
	var dfs func(x int)
	dfs = func(x int) {
		if len(graph[x]) == 0 {
			visited[x] = 1
			return
		}

		visited[x] = -1
		for _, y := range graph[x] {
			if visited[y] != 0 {
				if visited[y] < 0 {
					return
				}
				continue
			}
			dfs(y)
			if visited[y] < 0 {
				return
			}
		}

		visited[x] = 1
	}

	for i, v := range visited {
		if v == 0 {
			dfs(i)
		}
	}

	var ans []int
	for i, v := range visited {
		if v > 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
