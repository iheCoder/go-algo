package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var fast, slow *ListNode
	slow = head
	if slow == nil {
		return false
	}

	fast = head.Next
	if fast == nil {
		return false
	}
	for fast != slow {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return true
}
