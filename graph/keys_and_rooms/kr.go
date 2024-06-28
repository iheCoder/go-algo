package keys_and_rooms

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true

		for _, next := range rooms[i] {
			dfs(next)
		}
	}
	dfs(0)

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
