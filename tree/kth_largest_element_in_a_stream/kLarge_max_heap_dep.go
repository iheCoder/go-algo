package kth_largest_element_in_a_stream

import "sort"

// 最大堆明显是不行的
type intHeap []int

func (h *intHeap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h *intHeap) up(x int) {
	i := x
	var parent int
	for {
		parent = (i - 1) / 2
		if i == parent || h.Less(i, parent) {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func NewIntHeap(nums []int) *intHeap {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	h := (*intHeap)(&nums)
	//h.Init(nums)
	return h
}

func (h *intHeap) Init(nums []int) {
	for i := 0; i < len(nums); i++ {
		h.down(i)
	}
}

func (h *intHeap) down(x int) bool {
	i := x
	for {
		lc := i*2 + 1
		if lc < 0 || lc >= len(*h) {
			break
		}

		rc := lc + 1
		if rc < len(*h) && h.Less(i, rc) {
			h.Swap(i, lc)
			h.Swap(i, rc)
			i = rc
			continue
		}

		if h.Less(i, lc) {
			h.Swap(i, lc)
			i = lc
			continue
		}
		break
	}

	return i > x
}

func (h *intHeap) getTopK(k int) int {
	return (*h)[k-1]
}

func (h *intHeap) Less(i, j int) bool {
	if (*h)[i] < (*h)[j] {
		return true
	}
	return false
}

//type KthLargest struct {
//	k int
//	h *intHeap
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	h := NewIntHeap(nums)
//	return KthLargest{
//		k: k,
//		h: h,
//	}
//}
//
//func (this *KthLargest) Add(val int) int {
//	this.h.Push(val)
//	return this.h.getTopK(this.k)
//}
