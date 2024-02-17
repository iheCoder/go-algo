package remove_dup_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := make(map[int]bool)
	p := head
	dh := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dh
	for p != nil {
		if m[p.Val] {
			prev.Next = p.Next
		} else {
			prev = p
			m[p.Val] = true
		}
		p = p.Next
	}

	return dh.Next
}
