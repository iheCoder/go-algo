package kth_largest_element_in_a_stream

import "sort"

type linkedNode struct {
	val  int
	next *linkedNode
}

func NewLinkedNode(nums []int) *linkedNode {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	head := &linkedNode{}
	p := head
	for _, num := range nums {
		node := &linkedNode{
			val:  num,
			next: nil,
		}
		p.next = node
		p = node
	}
	return head
}

func (l *linkedNode) Push(x int) {
	p := l
	for p.next != nil && p.next.val >= x {
		p = p.next
	}

	node := &linkedNode{
		val:  x,
		next: p.next,
	}
	p.next = node
}

func (l *linkedNode) Get(target int) int {
	p := l
	for i := 0; i < target; i++ {
		p = p.next
	}
	return p.val
}

type KthLargest struct {
	k    int
	head *linkedNode
}

func Constructor(k int, nums []int) KthLargest {
	head := NewLinkedNode(nums)
	return KthLargest{
		k:    k,
		head: head,
	}
}

func (this *KthLargest) Add(val int) int {
	this.head.Push(val)
	return this.head.Get(this.k)
}
