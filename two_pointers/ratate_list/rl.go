package ratate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 1
	next := head.Next
	for next != nil {
		count++
		if next.Next == nil {
			next.Next = head
			break
		}
		next = next.Next
	}

	k = k % count
	r := count - k

	for i := 0; i < r; i++ {
		head = head.Next
		next = next.Next
	}
	next.Next = nil
	return head
}
