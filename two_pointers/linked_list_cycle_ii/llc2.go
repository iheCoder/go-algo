package linked_list_cycle_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	if fast == nil {
		return nil
	}

	slow = head
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
