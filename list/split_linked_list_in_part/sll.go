package split_linked_list_in_part

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	ans := make([]*ListNode, k)
	var n int
	p := head
	for p != nil {
		n++
		p = p.Next
	}

	c, rem := n/k, n%k
	prev := &ListNode{Next: head}
	for i := 0; i < k; i++ {
		expected := c
		if rem > 0 {
			rem--
			expected++
		}
		ans[i] = prev.Next
		p = prev
		for expected > 0 {
			prev = prev.Next
			expected--
		}
		p.Next = nil
	}
	return ans
}
