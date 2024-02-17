package odd_even_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 奇1指向奇2，偶1指向偶2
	oddTail, evenHead := head, head.Next
	p := head
	next := p.Next
	for next != nil {
		p.Next = next.Next
		if p.Next == nil {
			break
		}
		next.Next = p.Next.Next
		p = p.Next
		oddTail = p
		next = p.Next
	}
	oddTail.Next = evenHead
	return head
}
