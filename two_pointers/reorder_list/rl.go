package reorder_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mp := head
	slow, fast := head, head.Next
	for {
		mp = slow
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			mp = slow
			slow = slow.Next
			break
		}
	}

	prev, p := slow, slow.Next
	slow.Next = nil
	var next *ListNode
	for p != nil {
		next = p.Next
		p.Next = prev
		prev = p
		p = next
	}
	mp.Next = nil
	//mp.Next = prev

	i, j := head.Next, prev
	cur := head
	for j != nil {
		cur.Next = j
		j = j.Next
		cur = cur.Next
		cur.Next = i
		if i == nil {
			break
		}
		i = i.Next
		cur = cur.Next
	}
}
