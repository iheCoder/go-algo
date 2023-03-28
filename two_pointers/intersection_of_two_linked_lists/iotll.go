package intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var ic, jc int
	pa := headA
	pb := headB
	for {
		if pa == nil && pb == nil {
			break
		} else if pa == nil {
			for pb != nil {
				jc++
				pb = pb.Next
			}
			break
		} else if pb == nil {
			for pa != nil {
				ic++
				pa = pa.Next
			}
			break
		}
		pa = pa.Next
		pb = pb.Next
	}

	pa = headA
	pb = headB
	if ic > 0 {
		for ; ic > 0; ic-- {
			pa = pa.Next
		}
	} else if jc > 0 {
		for ; jc > 0; jc-- {
			pb = pb.Next
		}
	}

	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}

	return nil
}
