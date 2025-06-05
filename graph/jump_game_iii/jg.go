package jump_game_iii

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make(map[int]bool)
	reachZero := false
	var dfs func(i int)
	dfs = func(i int) {
		if i < 0 || i >= n {
			return
		}
		if reachZero {
			return
		}
		if arr[i] == 0 {
			reachZero = true
			return
		}

		if visited[i] {
			return
		}
		visited[i] = true

		l, r := i-arr[i], i+arr[i]
		dfs(l)
		dfs(r)
	}
	dfs(start)

	return reachZero
}
