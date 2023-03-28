package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next
	if fast == nil {
		return slow
	}

	for {
		slow = slow.Next
		if fast.Next == nil || fast.Next.Next == nil {
			return slow
		}
		fast = fast.Next.Next
	}
}
