package most_profit_assigning_work

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type pair struct {
		diff, p int
	}
	n := len(difficulty)
	ps := make([]*pair, 0, n)
	for i := 0; i < n; i++ {
		ps = append(ps, &pair{
			diff: difficulty[i],
			p:    profit[i],
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].diff < ps[j].diff || (ps[i].diff == ps[j].diff && ps[i].p < ps[j].p)
	})

	maxProfits := make([]int, n)
	var mx int
	for i := 0; i < n; i++ {
		if ps[i].p > mx {
			mx = ps[i].p
		}
		maxProfits[i] = mx
	}

	var ans int
	for i := 0; i < len(worker); i++ {
		x := sort.Search(n, func(j int) bool {
			return ps[j].diff > worker[i]
		})
		if x >= n {
			ans += maxProfits[n-1]
			//fmt.Println(maxProfits[n-1])
		} else if ps[x].diff == worker[i] {
			ans += maxProfits[x]
			//fmt.Println(maxProfits[x])
		} else if x > 0 {
			ans += maxProfits[x-1]
			//fmt.Println(maxProfits[x-1])
		} else {
			//fmt.Println(0)
		}
	}

	return ans
}
