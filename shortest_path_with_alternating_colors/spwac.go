package shortest_path_with_alternating_colors

// 这才是好的写法，没有考虑什么自环、环，但是却无比正确。我以往的思维需要一些新东西
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	g := [2][][]int{}
	for i := range g {
		g[i] = make([][]int, n)
	}
	for _, edge := range redEdges {
		g[0][edge[0]] = append(g[0][edge[0]], edge[1])
	}
	for _, e := range blueEdges {
		g[1][e[0]] = append(g[1][e[0]], e[1])
	}

	type pair struct {
		// 节点index
		i int
		// color
		c int
	}
	q := []pair{
		{
			i: 0,
			c: 0,
		},
		{
			i: 0,
			c: 1,
		},
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	// 记录访问过的点和颜色
	vis := make([][2]bool, n)

	d := 0
	// 遍历队列，计算最短距离
	// 遍历的时候就是判断有没有访问过，没有访问过就可以接着走，距离就是根据从到0点距离计算出来的
	// 但我唯一的问题是这真的是最短的吗
	for len(q) > 0 {
		for k := len(q); k > 0; k-- {
			p := q[0]
			q = q[1:]
			i, c := p.i, p.c
			//
			if ans[i] == -1 {
				ans[i] = d
			}
			// 标记点i颜色c访问过了
			vis[i][c] = true
			// 取目标颜色
			c ^= 1
			// 遍历图，若没有访问过，就加入到队列中
			for _, j := range g[c][i] {
				if !vis[j][c] {
					q = append(q, pair{
						i: j,
						c: c,
					})
				}
			}
		}
		d++
	}
	return ans
}
