package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	if head == nil {
		return head
	}

	i := 1
	p := head
	var prev, next *ListNode
	var l, lp *ListNode
	for p != nil {
		if i == left {
			l = p
			lp = prev
			break
		}
		i++
		prev = p
		p = p.Next
	}

	for p != nil {
		next = p.Next
		p.Next = prev
		if i == right {
			l.Next = next
			if lp != nil {
				lp.Next = p
			} else {
				head = p
			}
			break
		}
		i++
		prev = p
		p = next
	}
	return head
}
