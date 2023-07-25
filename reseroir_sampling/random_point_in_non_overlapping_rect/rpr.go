package random_point_in_non_overlapping_rect

import (
	"math/rand"
	"sort"
)

type Solution struct {
	rects [][]int
	sums  []int
}

func Constructor(rects [][]int) Solution {
	sums := make([]int, len(rects)+1)
	var sum int
	for i, rect := range rects {
		x0, y0, x1, y1 := rect[0], rect[1], rect[2], rect[3]
		sum += (x1 - x0 + 1) * (y1 - y0 + 1)
		sums[i+1] = sum
	}
	return Solution{
		rects: rects,
		sums:  sums,
	}
}

func (this *Solution) Pick() []int {
	k := rand.Intn(this.sums[len(this.sums)-1])
	// k要加一是为了找到恰好大于k的值
	// 在极端情况下，k等于最大值的时候i便会指向最后一个
	i := sort.SearchInts(this.sums, k+1) - 1
	rect := this.rects[i]
	x0, y0, y1 := rect[0], rect[1], rect[3]
	d := k - this.sums[i]
	// 点(x,y)的值为x*(y1-y0+1)+y
	dx := d / (y1 - y0 + 1)
	dy := d % (y1 - y0 + 1)
	return []int{x0 + dx, y0 + dy}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
