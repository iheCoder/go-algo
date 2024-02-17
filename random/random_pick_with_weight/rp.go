package random_pick_with_weight

import (
	"math/rand"
	"sort"
)

type Solution struct {
	sums  []int
	total int
}

func Constructor(w []int) Solution {
	sums := make([]int, len(w)+1)
	for i, v := range w {
		sums[i+1] += v + sums[i]
	}
	return Solution{sums: sums, total: sums[len(sums)-1]}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.total) + 1
	return sort.SearchInts(this.sums, x) - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
