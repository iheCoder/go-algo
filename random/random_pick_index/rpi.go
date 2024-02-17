package random_pick_index

import "math/rand"

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i, num := range nums {
		if len(m[num]) == 0 {
			m[num] = make([]int, 0)
		}
		m[num] = append(m[num], i)
	}
	return Solution{m: m}
}

func (this *Solution) Pick(target int) int {
	if len(this.m[target]) == 1 {
		return this.m[target][0]
	} else {
		l := len(this.m[target])
		x := rand.Intn(l)
		return this.m[target][x]
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
