package redundant_connection

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}

	var union func(i, j int) bool
	union = func(i, j int) bool {
		ip, jp := find(i), find(j)
		if ip == jp {
			return false
		}

		parents[jp] = ip
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}
