package rm_zero_sum_consecutive_nodes_from_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	dh := &ListNode{Next: head}
	prev := dh
	var cnt int
	for prev != nil {
		cur := prev.Next
		cnt = 0
		for cur != nil {
			cnt += cur.Val
			if cnt == 0 {
				prev.Next = cur.Next
			}
			cur = cur.Next
		}
		prev = prev.Next
	}

	return dh.Next
}

// 	arr := make([]int, 0)
//	sums := make([]int, 1)
//	p := head
//	for p != nil {
//		arr = append(arr, p.Val)
//		sums = append(sums, p.Val+sums[len(sums)-1])
//	}
//
//	n := len(arr)
//	for _, sum := range sums {
//
//	}
