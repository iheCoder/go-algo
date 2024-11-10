package shortest_path_with_alternating_colors

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	rm, bm := make(map[int][]int), make(map[int][]int)
	for _, edge := range redEdges {
		rm[edge[0]] = append(rm[edge[0]], edge[1])
	}
	for _, edge := range blueEdges {
		bm[edge[0]] = append(bm[edge[0]], edge[1])
	}

	type state struct {
		node, color, steps int
	}

	queue := []state{{0, 1, 0}, {0, -1, 0}}
	visited := make(map[int]bool)
	visited[1] = true
	visited[-1] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if ans[p.node] == -1 || ans[p.node] > p.steps {
			ans[p.node] = p.steps
		}

		nextEdges := rm
		if p.color == -1 {
			nextEdges = bm
		}

		for _, next := range nextEdges[p.node] {
			newState := state{
				node:  next,
				color: -p.color,
				steps: p.steps + 1,
			}
			key := (newState.node + 1) * newState.color

			if visited[key] {
				continue
			}

			queue = append(queue, newState)
			visited[key] = true
		}
	}

	return ans
}

func minInt(x, y int) int {
	if x < 0 {
		return y
	}

	if x < y {
		return x
	}

	return y
}

// 	fromNum := 10000
//	visited := make(map[int]bool)
//	var f func(at, color, sum int)
//	f = func(at, color, sum int) {
//		ans[at] = minInt(ans[at], sum)
//
//		em := rm
//		if color == -1 {
//			em = bm
//		}
//
//		for _, to := range em[at] {
//			key := -color * (at*fromNum + to)
//			if visited[key] {
//				continue
//			}
//
//			visited[key] = true
//			f(to, -color, sum+1)
//
//			delete(visited, key)
//		}
//	}
//
//	f(0, 1, 0)
//	f(0, -1, 0)
//
//	return ans
