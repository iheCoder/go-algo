package sort_int_by_power_value

import (
	"sort"
)

func getKth(lo int, hi int, k int) int {
	memo := map[int]int{
		1: 0,
	}

	getWidth := func(x int) int {
		if c, ok := memo[x]; ok {
			return c
		}

		original := x
		var cnt int
		defer func() {
			memo[x] = cnt
		}()

		for x > 1 {
			if x%2 == 1 {
				x = 3*x + 1
			} else {
				x /= 2
			}

			cnt++

			if c, ok := memo[x]; ok {
				cnt += c
				break
			}
		}

		memo[original] = cnt
		return cnt
	}

	type pair struct {
		x, weight int
	}
	ps := make([]pair, 0)

	for i := lo; i <= hi; i++ {
		ps = append(ps, pair{
			x:      i,
			weight: getWidth(i),
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].weight < ps[j].weight || (ps[i].weight == ps[j].weight && ps[i].x < ps[j].x)
	})

	return ps[k-1].x
}
