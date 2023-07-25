package min_height_trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	ad := make([][]int, n)
	for i := 0; i < n; i++ {
		ad[i] = make([]int, 0)
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		ad[x] = append(ad[x], y)
		ad[y] = append(ad[y], x)
	}

	en := make([]int, n)
	q := make([]int, 0)
	for i, d := range ad {
		en[i] = len(d)
		if en[i] == 1 {
			q = append(q, i)
		}
	}

	rem := n
	var vc int
	for rem > 2 {
		vc = len(q)
		rem -= vc
		for _, x := range q {
			for _, y := range ad[x] {
				en[y]--
				if en[y] == 1 {
					q = append(q, y)
				}
			}
		}
		q = q[vc:]
	}
	return q
}
