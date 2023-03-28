package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	large := &ListNode{}
	lh := large
	sh := small

	p := head
	for p != nil {
		if p.Val < x {
			small.Next = p
			small = small.Next
		} else {
			large.Next = p
			large = large.Next
		}
		p = p.Next
	}

	small.Next = lh.Next
	large.Next = nil
	return sh.Next
}
