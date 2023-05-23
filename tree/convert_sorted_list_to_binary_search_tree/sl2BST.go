package convert_sorted_list_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	prev := head
	slow, fast := head, head.Next
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}

	prev.Next = nil
	left := sortedListToBST(head)
	right := sortedListToBST(slow.Next)
	root := &TreeNode{
		Val:   slow.Val,
		Left:  left,
		Right: right,
	}
	return root
}

// 	prev := root
//	p := head.Next
//	for p != nil {
//		cur := &TreeNode{Val: p.Val}
//		switch {
//		case prev.Left == nil && prev.Right == nil:
//			cur.Left = prev
//			prev = cur
//			root = cur
//		case prev.Right == nil:
//			prev.Right = cur
//		default:
//			cur.Left = prev
//			root = cur
//			prev = root
//		}
//	}
//
//	return root
//  翻转可行乎？
