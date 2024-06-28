package all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	var ans [][]int
	n := len(graph)
	var dfs func(trace []int, cur int)
	dfs = func(trace []int, cur int) {
		if cur == n-1 {
			ans = append(ans, trace)
			return
		}
		for _, next := range graph[cur] {
			ct := make([]int, len(trace))
			copy(ct, trace)
			ct = append(ct, next)
			dfs(ct, next)
		}
	}
	dfs([]int{0}, 0)

	return ans
}
