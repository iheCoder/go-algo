package open_the_lock

import "container/heap"

type astar struct {
	// g到当前点的实际距离，h到目标点的启发距离
	g, h   int
	status string
}
type hp []astar

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].g+h[i].h < h[j].g+h[j].h }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(astar)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getH(status, target string) int {
	var ans int
	for i := 0; i < 4; i++ {
		dist := abs(int(status[i]) - int(target[i]))
		ans += minInt(dist, 10-dist)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func openLock(deadends []string, target string) int {
	const source = "0000"
	if target == source {
		return 0
	}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[source] {
		return -1
	}

	// 得到status转动一位的所有可能
	get := func(status string) []string {
		s := []byte(status)
		var ret []string
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))

			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))

			s[i] = b
		}
		return ret
	}

	h := hp{{0, getH(source, target), source}}
	seen := map[string]bool{source: true}
	for len(h) > 0 {
		node := heap.Pop(&h).(astar)
		for _, next := range get(node.status) {
			if !seen[next] && !dead[next] {
				if next == target {
					return node.g + 1
				}
				seen[next] = true
				heap.Push(&h, astar{
					g:      node.g + 1,
					h:      getH(next, target),
					status: next,
				})
			}
		}
	}
	return -1
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 	dm := make(map[string]bool)
//	for _, deadend := range deadends {
//		dm[deadend] = true
//	}
//
//	visited := make(map[string]bool)
//	getNextByte := func(b byte) byte {
//		if b == '9' {
//			return '0'
//		}
//		return b + 1
//	}
//	getPrevByte := func(b byte) byte {
//		if b == '0' {
//			return '9'
//		}
//		return b - 1
//	}
//
//	var dfs func(s string, d int)
//	ans := math.MaxInt
//	dfs = func(s string, d int) {
//		if s == target {
//			ans = minInt(ans, d)
//			return
//		}
//
//		visited[s] = true
//		cps := make([]byte, len(s))
//		copy(cps, s)
//		var x string
//		for i := 0; i < len(s); i++ {
//			sp := getPrevByte(s[i])
//			cps[i] = sp
//			x = string(cps)
//
//			if !visited[x] && !dm[x] {
//				dfs(x, d+1)
//			}
//
//			sn := getNextByte(s[i])
//			cps[i] = sn
//			x = string(cps)
//			if !visited[x] && !dm[x] {
//				dfs(x, d+1)
//			}
//			cps[i] = s[i]
//		}
//	}
//	dfs("0000", 0)
//	if ans == math.MaxInt {
//		return -1
//	}
//	return ans
