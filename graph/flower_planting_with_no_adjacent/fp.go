package flower_planting_with_no_adjacent

func gardenNoAdj(n int, paths [][]int) []int {
	passTo := make([][]int, n+1)
	cantorPair := func(a, b int) int {
		if a > b {
			a, b = b, a
		}
		return (a+b)*(a+b+1)/2 + b
	}
	visited := make(map[int]bool)
	for _, path := range paths {
		x, y := path[0], path[1]
		key := cantorPair(x, y)
		if visited[key] {
			continue
		}
		passTo[x] = append(passTo[x], y)
		passTo[y] = append(passTo[y], x)
		visited[key] = true
	}

	type pair struct {
		adjColor int
		index    int
	}
	ps := make([]pair, n+1)

	getRemColor := func(num uint8) int {
		if num == 0 {
			return 1
		}

		for i := 0; i < 4; i++ {
			if num&(1<<i) == 0 {
				return i + 1
			}
		}

		return 1
	}

	p := 1
	var selectBits uint8
	for p <= n {
		selectBits = 0
		for _, adj := range passTo[p] {
			if ps[adj].adjColor == 0 {
				continue
			}

			selectBits |= 1 << (ps[adj].adjColor - 1)
		}

		ps[p].adjColor = getRemColor(selectBits)
		p++
	}

	ans := make([]int, n)
	for i := 1; i <= n; i++ {
		ans[i-1] = ps[i].adjColor
	}
	return ans
}

// it's seems stupid
// 	getRemColor := func(count, prod int) int {
//		if count == 0 {
//			return 1
//		}
//
//		// 1,1,2 2 3,4
//		// 1,2,2 4 3,4
//		// 1,1,3 3 2,4
//		// 1,3,3 9 2,4
//		// 1,1,4 4 2,3
//		// 1,4,4 16 2,3
//		// 2,2,3 12 1,4
//		// 2,3,3 18 1,4
//		// 2,2,4 16 1,3
//		// 2,4,4 32 1,3
//		// 3,3,4 36 1,2
//		// 3,4,4 48 1,2
//		if count == 3 {
//			return 24 / prod
//		}
//		if count == 1 {
//			return prod%4 + 1
//		}
//
//		// 1 2,3,4
//		// 2 3,4
//		// 4 1,3,4
//		// 6 1,4
//		// 8 1,3
//		// 9 1,2,4
//		// 12 1,2
//		// 16 1,2,3
//
//		if prod < 6 {
//			return 3
//		} else if prod > 8 {
//			return 2
//		}
//
//		return 1
//	}
