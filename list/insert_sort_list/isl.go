package insert_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dh := &ListNode{Next: head}
	lastSorted, cur := head, head.Next
	for cur != nil {
		if cur.Val >= lastSorted.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dh
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}

			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return dh.Next
}
