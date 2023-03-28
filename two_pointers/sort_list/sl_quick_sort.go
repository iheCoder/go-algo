package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortListQuickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, _ := mergeQS(head, nil)
	return l
}

func mergeQS(head, tail *ListNode) (*ListNode, *ListNode) {
	if head == tail {
		return head, tail
	}
	l, t := quickSort(head, tail)
	if l.Next == t {
		return l, t
	}

	l1, t1 := mergeQS(l, head)
	l2, t2 := mergeQS(head.Next, t)
	t1.Next = l2
	return l1, t2
}

func quickSort(head, tail *ListNode) (*ListNode, *ListNode) {
	p := head
	l := head
	next := p
	newTail := head
	for p != nil && p != tail {
		next = next.Next
		if p.Val < head.Val {
			// 若左边界指向p，在左边界的指向还没有改变之前，让p又指向了左边界，那么就会形成环
			// !!! 错误，也不是左边界，而是新边界。新边界指向没有遍历过的位置
			newTail.Next = next
			p.Next = l
			l = p
		} else {
			newTail = p
		}
		p = next
	}
	newTail.Next = tail

	return l, tail
}
