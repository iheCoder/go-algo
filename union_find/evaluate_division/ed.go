package evaluate_division

type unionFind struct {
	parents []int
	// 指向父节点的权重
	weights []float64
}

func newUnionFind(n int) *unionFind {
	parents := make([]int, n)
	weights := make([]float64, n)
	// 初始设置所有父节点为自己，权重为1
	for i := 0; i < n; i++ {
		parents[i] = i
		weights[i] = 1.0
	}
	return &unionFind{
		parents: parents,
		weights: weights,
	}
}

func (u *unionFind) find(x int) int {
	if u.parents[x] == x {
		return x
	}

	original := u.parents[x]
	u.parents[x] = u.find(original)
	u.weights[x] *= u.weights[original]
	return u.parents[x]
}

func (u *unionFind) union(x, y int, v float64) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}

	// 应该让除数指向被除数
	u.parents[rootX] = rootY
	u.weights[rootX] = u.weights[y] * v / u.weights[x]
}

func (u *unionFind) isConnected(x, y int) float64 {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return u.weights[x] / u.weights[y]
	}
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	uf := newUnionFind(2 * n)
	m := make(map[string]int)
	var id int
	var v1, v2 string
	for i := 0; i < n; i++ {
		v1, v2 = equations[i][0], equations[i][1]
		if _, ok := m[v1]; !ok {
			m[v1] = id
			id++
		}
		if _, ok := m[v2]; !ok {
			m[v2] = id
			id++
		}

		uf.union(m[v1], m[v2], values[i])
	}

	ans := make([]float64, 0, len(queries))
	var ok0, ok1 bool
	var x, y int
	for _, query := range queries {
		x, ok0 = m[query[0]]
		y, ok1 = m[query[1]]
		if !ok0 || !ok1 {
			ans = append(ans, -1.0)
			continue
		}

		ans = append(ans, uf.isConnected(x, y))
	}
	return ans
}
