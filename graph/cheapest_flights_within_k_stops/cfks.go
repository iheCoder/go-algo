package cheapest_flights_within_k_stops

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1 << 31
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	dp[src] = 0

	ans := inf
	for t := 0; t <= k; t++ {
		g := make([]int, n)
		for i := 0; i < n; i++ {
			g[i] = inf
		}

		for _, flight := range flights {
			i, j, p := flight[0], flight[1], flight[2]
			g[j] = minInt(g[j], dp[i]+p)
		}

		dp = g
		ans = minInt(ans, dp[dst])
	}

	if ans == inf {
		return -1
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

// 	dist := make([]int, n)
//	const maxVal = 1 << 31
//	for i := 0; i < n; i++ {
//		dist[i] = maxVal
//	}
//	dist[src] = 0
//
//	g := make([][]int, n)
//	for i := 0; i < n; i++ {
//		g[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			g[i][j] = maxVal
//		}
//	}
//	for _, flight := range flights {
//		g[flight[0]][flight[1]] = flight[2]
//	}
//
//	k++
//	visited := make([]bool, n)
//	transits := make([]int, n)
//	for i := 0; i < n; i++ {
//		x := -1
//		for y, u := range visited {
//			if !u && (x == -1 || dist[x] > dist[y]) {
//				x = y
//			}
//		}
//
//		visited[x] = true
//		for y, p := range g[x] {
//			zp := dist[x] + p
//			zt := transits[x]
//			// 通往2不可能比通往1小
//			// 最短路径本来就不关系k的
//			if zp < dist[y] && zt < k {
//				transits[y] = zt + 1
//				dist[y] = zp
//			}
//		}
//	}
//
//	if transits[dst] > k || dist[dst] == maxVal {
//		return -1
//	}
//
//	return dist[dst]
