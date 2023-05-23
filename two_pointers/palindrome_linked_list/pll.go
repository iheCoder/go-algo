package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var fast, slow *ListNode
	slow = head
	fast = head.Next
	// 该如何使得在偶数时，slow指向偏右的那个元素呢？
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	prev := slow
	cur := slow.Next
	slow.Next = nil
	next := cur
	for next != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	p1 := head
	p2 := prev
	for p1 != p2 && p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
