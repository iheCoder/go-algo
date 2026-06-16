package shortest_path_visiting_all_nodes

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 1 {
		return 0
	}

	fullMask := (1 << n) - 1

	type State struct {
		node int
		mask int
		dist int
	}
	queue := make([]State, 0, n)
	for i, _ := range graph {
		queue = append(queue, State{
			node: i,
			mask: 1 << i,
		})
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, 1<<n)
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.mask == fullMask {
			return cur.dist
		}

		for _, next := range graph[cur.node] {
			nextMask := cur.mask | 1<<next

			if visited[next][nextMask] {
				continue
			}
			visited[next][nextMask] = true

			queue = append(queue, State{
				node: next,
				mask: nextMask,
				dist: cur.dist + 1,
			})
		}
	}

	return -1
}
