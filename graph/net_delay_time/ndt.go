package net_delay_time

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	if len(times) < n-1 {
		return -1
	}

	const inf = math.MaxInt64 / 2
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[k-1] = 0

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = inf
		}
	}
	for _, td := range times {
		g[td[0]-1][td[1]-1] = td[2]
	}

	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		// 找到未访问的最短路径的节点x
		for y, u := range visited {
			if !u && (x == -1 || dist[x] > dist[y]) {
				x = y
			}
		}

		visited[x] = true
		// 从节点x找到其通往的节点，然后求最短路径
		for y, w := range g[x] {
			dist[y] = minInt(dist[y], dist[x]+w)
		}
	}

	var ans int
	for _, p := range dist {
		if p == inf {
			return -1
		}
		ans = maxInt(ans, p)
	}
	return ans
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 	visited := make([]int, n+1)
//	for i := 0; i <= n; i++ {
//		visited[i] = math.MaxInt
//	}
//
//	m := make(map[int][]*pair)
//	for _, tm := range times {
//		m[tm[0]] = append(m[tm[0]], &pair{
//			target: tm[1],
//			weight: tm[2],
//		})
//	}
//
//	var dfs func(i, b int) int
//	var x int
//	dfs = func(i, b int) int {
//		visited[i] = b
//		a := b
//		x = maxInt(x, b)
//		for _, p := range m[i] {
//			if visited[p.target] <= b+p.weight {
//				continue
//			}
//			a = maxInt(a, dfs(p.target, b+p.weight))
//		}
//		return a
//	}
//
//	dfs(k, 0)
//	for i := 1; i < len(visited); i++ {
//		if visited[i] == math.MaxInt {
//			return -1
//		}
//	}
//	return x
