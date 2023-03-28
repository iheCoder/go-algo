package sort_list

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var count int
	p := head
	for p != nil {
		count++
		p = p.Next
	}

	for i := 1; i < count; i <<= 1 {
		head = divide(head, i)
	}

	return head
}

func divide(h *ListNode, size int) *ListNode {
	dh := &ListNode{
		Val:  0,
		Next: h,
	}
	next := h
	cur := dh.Next
	prev := dh
	tail := dh
	for next != nil {
		h1 := cur
		for i := 0; i < size && cur != nil; i++ {
			prev = cur
			cur = cur.Next
		}

		prev.Next = nil
		h2 := cur
		for i := 0; i < size && cur != nil; i++ {
			prev = cur
			cur = cur.Next
		}
		if cur == nil {
			next = nil
		} else {
			next = cur
			prev.Next = nil
		}

		h3 := mergeSort(h1, h2)
		tail.Next = h3
		temp := h3
		for temp.Next != nil {
			temp = temp.Next
		}
		tail = temp

		cur = next
	}
	return dh.Next
}

func mergeSort(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	} else if h2 == nil {
		return h1
	}

	vh := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := vh
	hp1 := h1
	hp2 := h2
	for {
		if hp1 == nil {
			p.Next = hp2
			break
		} else if hp2 == nil {
			p.Next = hp1
			break
		}

		if hp1.Val < hp2.Val {
			p.Next = hp1
			hp1 = hp1.Next
		} else {
			p.Next = hp2
			hp2 = hp2.Next
		}
		p = p.Next
	}

	return vh.Next
}
