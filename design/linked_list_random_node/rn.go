package linked_list_random_node

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	var ans int
	for n, i := this.head, 1; n != nil; n = n.Next {
		if rand.Intn(i) == 0 {
			ans = n.Val
		}
		i++
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
