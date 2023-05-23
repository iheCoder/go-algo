package kth_largest_element_in_a_stream

import "testing"

func TestIntHeap(t *testing.T) {
	nums := []int{3, 5, 1, 9, 2}
	h := NewIntHeap(nums)
	t.Log(*h)
	t.Log(h.getTopK(3))
	h.Push(10)
	t.Log(h.getTopK(3))
	h.Push(0)
	t.Log(h.getTopK(3))
}

func TestLinkedNode(t *testing.T) {
	nums := []int{3, 5, 1, 9, 2}
	head := NewLinkedNode(nums)
	t.Log(head.Get(3))
	head.Push(10)
	t.Log(head.Get(3))
	head.Push(0)
	t.Log(head.Get(3))
}
