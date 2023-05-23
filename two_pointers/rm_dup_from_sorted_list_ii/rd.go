package rm_dup_from_sorted_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dh := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dh
	cur := dh.Next
	next := cur.Next

	isDup := false
	for next != nil {
		if cur.Val != next.Val {
			if !isDup {
				prev.Next = cur
				prev = cur
			}
			isDup = false
		} else {
			isDup = true
		}
		cur = next
		next = cur.Next
	}
	if cur != nil {
		if !isDup {
			prev.Next = cur
		} else {
			prev.Next = nil
		}
	}

	return dh.Next
}
