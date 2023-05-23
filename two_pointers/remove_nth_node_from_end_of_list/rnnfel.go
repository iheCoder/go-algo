package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	i := 1
	for fast != nil && i < n {
		i++
		fast = fast.Next
	}

	fast = fast.Next
	slow := head
	vh := &ListNode{
		Val:  0,
		Next: slow,
	}
	prev := vh
	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}
	prev.Next = slow.Next
	slow.Next = nil

	return vh.Next
}
